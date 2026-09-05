package lamportsmutualexclusionalgorithm

type MessageType int

const (
	Request MessageType = iota
	Release
	Other
)

type Message struct {
	Clock int
	T     MessageType
	// Content []byte
}

func (m Message) CompareTo(other Message) int {
	return m.Clock - other.Clock
}
