package lamportsmutualexclusionalgorithm

import (
	"fmt"
	"math/rand/v2"
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
	msg, ok := <-p.Me
	if !ok {
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
			fmt.Println("ERROR STATE: RECIEVED RELEASE MSG, BUT NO PROCESSES IN QUEUE")
		}
		if msg.Pid != popMsg.Pid {
			// error state
			fmt.Println("ERROR STATE: RECIEVED RELEASE MSG FROM WRONG PROCESS")
		}
	}
	if msg.T == Acknowledge {
		p.AckCount += 1
	}
}

func (p *Process) SendMessage(msg Message, peer chan<- Message) {
	p.Clock += 1
	peer <- msg
}

func (p *Process) RequestCS() {
	p.Clock += 1
	p.AckCount = 0
	p.CSRequested = true
	msg := Message{p.Pid, p.Clock, Request}
	for _, peer := range p.Peers {
		p.SendMessage(msg, peer)
	}
}

func (p *Process) Run() {
	for {
		// p.InternalEvent()
		p.AttemptReceiveMessage()

		if p.CSRequested {
			if msg, ok := p.CSPriority.Peek(); ok && msg.Pid == p.Pid && p.AckCount == p.N {
				// enter CS
				// do some computation

				time.Sleep(1 * time.Second)

				// release CS
				for _, peer := range p.Peers {
					p.SendMessage(Message{
						p.Pid,
						p.Clock,
						Release,
					}, peer)
				}
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
