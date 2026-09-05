package lamportsmutualexclusionalgorithm

type Process struct {
	clock int

	receiver    <-chan Message
	senders     []chan<- Message
	ackPriority MessageCollection
}
