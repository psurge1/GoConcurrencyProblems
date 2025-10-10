package pcproblem

import (
	"fmt"
	"strings"
	"sync"
)

type Data struct {
	id int
}

type PCQueue struct {
	lock     sync.Mutex
	hasSpace *sync.Cond
	hasData  *sync.Cond
	buffer   []Data
	capacity int
	size     int
	p        int
	c        int
	dataId   int
}

func formatReturnString(tp string, buffer []Data, item Data) string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("\n%s data: %v\n", tp, item))
	sb.WriteString(fmt.Sprintf("Queue After %s:\n", tp))
	for _, el := range buffer {
		sb.WriteString(fmt.Sprintf("%d ", el.id))
	}
	sb.WriteString("\n")
	return sb.String()
}

func CreatePCQueue(bufferWidth int) *PCQueue {
	pcq := &PCQueue{
		buffer:   make([]Data, bufferWidth),
		capacity: bufferWidth,
		size:     0,
		p:        0,
		c:        0,
		dataId:   1,
	}

	pcq.hasSpace = sync.NewCond(&pcq.lock)
	pcq.hasData = sync.NewCond(&pcq.lock)

	return pcq
}

func (buffer *PCQueue) produce() error {
	buffer.lock.Lock()
	defer buffer.lock.Unlock()

	for buffer.size >= buffer.capacity {
		buffer.hasSpace.Wait()
	}

	buffer.buffer[buffer.p] = Data{buffer.dataId}
	data := buffer.buffer[buffer.p]

	buffer.dataId++
	buffer.p = (buffer.p + 1) % buffer.capacity
	buffer.size++

	displayStr := formatReturnString("Produce", buffer.buffer, data)
	// TODO: Implement channel so IO isn't called inside critical section
	fmt.Println(displayStr)

	buffer.hasData.Signal()

	return nil
}

func (buffer *PCQueue) consume() error {
	buffer.lock.Lock()
	defer buffer.lock.Unlock()

	for buffer.size <= 0 {
		buffer.hasData.Wait()
	}

	data := buffer.buffer[buffer.c]

	buffer.buffer[buffer.c].id = 0
	buffer.c = (buffer.c + 1) % buffer.capacity
	buffer.size--

	displayStr := formatReturnString("Consume", buffer.buffer, data)
	// TODO: Implement channel so IO isn't called inside critical section
	fmt.Println(displayStr)

	buffer.hasSpace.Signal()

	return nil
}
