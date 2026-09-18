package lamportsmutualexclusionalgorithm

import (
	"math/rand/v2"
	"strconv"
	"sync/atomic"
	"time"
)

type Process struct {
	Pid   int
	Clock int

	Me         <-chan Message
	Peers      []chan<- Message
	CSPriority MessageHeap

	N           int
	AckCount    int
	CSRequested bool

	Logs chan string
}

func InitSystem(N int, logs chan string) []*Process {
	processes := make([]*Process, N)
	channels := make([]chan<- Message, N)
	for i := range N {
		pChannel := make(chan Message, 1000)
		processes[i] = &Process{
			i,
			1,
			pChannel,
			make([]chan<- Message, N),
			NewMessageHeap(),
			N,
			0,
			false,
			logs,
		}
		channels[i] = pChannel
	}
	for i := range N {
		copy(processes[i].Peers, channels)
	}
	return processes
}

func (p *Process) Init() {
	p.Clock = 1
	p.Me = make(<-chan Message, 1000)

	p.AckCount = 0
	p.CSRequested = false
}

func (p *Process) InternalEvent() {
	p.Clock += 1
}

func (p *Process) AttemptReceiveMessage() {
	msg := Message{}
	msgReceived := false
	select {
	case msg = <-p.Me:
		msgReceived = true
	default:
	}
	if !msgReceived {
		return
	}

	p.Clock = max(p.Clock, msg.Clock) + 1

	if msg.T == Request {
		p.CSPriority.Push(msg)
		p.SendMessage(Message{
			p.Pid,
			p.Clock,
			Acknowledge,
		}, p.Peers[msg.Pid])
	}
	if msg.T == Release {
		popMsg, ok := p.CSPriority.Pop()
		if !ok {
			p.Logs <- "ERROR STATE: RECIEVED RELEASE MSG, BUT NO PROCESSES IN QUEUE"
		}
		if msg.Pid != popMsg.Pid {
			// error state
			p.Logs <- "ERROR STATE: RECIEVED RELEASE MSG FROM WRONG PROCESS"
		}
	}
	if msg.T == Acknowledge {
		p.AckCount += 1
	}
}

func (p *Process) SendMessage(msg Message, peer chan<- Message) {
	peer <- msg
}

func (p *Process) RequestCS() {
	p.Clock += 1
	p.AckCount = 0
	p.CSRequested = true
	p.Logs <- "Process " + strconv.Itoa(p.Pid) + " REQUESTING CS at " + strconv.Itoa(p.Clock)
	msg := Message{p.Pid, p.Clock, Request}
	for _, peer := range p.Peers {
		p.SendMessage(msg, peer)
	}
}

func (p *Process) ReleaseCS() {
	p.Clock += 1
	p.CSPriority.Pop()
	p.CSRequested = false
	p.AckCount = 0
	p.Logs <- "Process " + strconv.Itoa(p.Pid) + " RELEASING CS at " + strconv.Itoa(p.Clock)
	for idx, peer := range p.Peers {
		if idx != p.Pid {
			p.SendMessage(Message{
				p.Pid,
				p.Clock,
				Release,
			}, peer)
		}
	}
}

func (p *Process) Run() {
	numProcessesInCS := atomic.Int32{}
	for {
		// p.InternalEvent()
		p.AttemptReceiveMessage()

		if p.CSRequested {
			if msg, ok := p.CSPriority.Peek(); ok && msg.Pid == p.Pid && p.AckCount == p.N {
				// enter CS
				p.Logs <- "Process " + strconv.Itoa(p.Pid) + " ENTERING CS, original request timestamp: " + strconv.Itoa(msg.Clock)

				// ensure only one process exists in the critical section
				if numProcessesInCS.Add(1) > 1 {
					p.Logs <- "TOO MANY PROCESSES IN CS"
				}
				// do some computation in the CS
				time.Sleep(1 * time.Second)
				numProcessesInCS.Add(-1)

				// release CS
				p.ReleaseCS()
			}
		} else {
			csRequestChance := 20 // percentage
			probability := rand.IntN(101)
			if probability < csRequestChance {
				p.RequestCS()
			}
		}
	}
}
