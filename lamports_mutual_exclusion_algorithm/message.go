package lamportsmutualexclusionalgorithm

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
