package lamportsmutualexclusionalgorithm

type Process struct {
	Clock int

	Me          <-chan Message
	Peers       []chan<- Message
	AckPriority MessageHeap
}

func (p *Process) InternalEvent() {
	p.Clock += 1
}

func (p *Process) ReceiveMessage(msg Message) {
	p.Clock = max(p.Clock, msg.Clock) + 1

	// order these if statements by frequency to optimize for the branch predictor
	if msg.T == Request {
	}
	if msg.T == Release {
	}
	if msg.T == Other {
	}
}

func (p *Process) SendMessage(msg Message, peer chan<- Message) {
	peer <- msg
}

func (p *Process) RequestCS() {
	p.Clock += 1
	msg := Message{p.Clock, Request}
	for _, peer := range p.Peers {
		peer <- msg
	}
}
