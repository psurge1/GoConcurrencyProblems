package ricartagrawalamutualexclusionalgorithm

type MessageType int

const (
	Request MessageType = iota
	Release
	Acknowledge
	Other
)

type Message struct {
	Pid   int
	Clock int
	T     MessageType
	// Content []byte
}

func (m Message) CompareTo(other Message) int {
	clockDiff := m.Clock - other.Clock
	if clockDiff != 0 {
		return clockDiff
	}
	return m.Pid - other.Pid
}

type Process struct {
	Pid   int
	Clock int

	Me    <-chan Message
	Peers []chan<- Message

	N               int
	PermissionCount int
	CSRequested     bool

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
