package pcproblem

import (
	"fmt"
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
	var data Data
	var queueCopy []Data
	defer func() {
		fmt.Printf("\nProduced data: %v\n", data)
		fmt.Printf("Queue After Production:\n")
		for _, el := range queueCopy {
			fmt.Printf("%d ", el.id)
		}
		fmt.Println()
	}()

	buffer.lock.Lock()
	defer buffer.lock.Unlock()

	for buffer.size >= buffer.capacity {
		buffer.hasSpace.Wait()
	}

	buffer.buffer[buffer.p] = Data{buffer.dataId}
	data = buffer.buffer[buffer.p]

	buffer.dataId++
	buffer.p = (buffer.p + 1) % buffer.capacity
	buffer.size++

	queueCopy = append([]Data(nil), buffer.buffer...)

	buffer.hasData.Signal()

	return nil
}

func (buffer *PCQueue) consume() error {
	var data Data
	var queueCopy []Data
	defer func() {
		fmt.Printf("\nConsumed data: %v\n", data)
		fmt.Printf("Queue After Consumption:\n")
		for _, el := range queueCopy {
			fmt.Printf("%d ", el.id)
		}
		fmt.Println()
	}()

	buffer.lock.Lock()
	defer buffer.lock.Unlock()

	for buffer.size <= 0 {
		buffer.hasData.Wait()
	}

	data = buffer.buffer[buffer.c]

	buffer.buffer[buffer.c].id = 0
	buffer.c = (buffer.c + 1) % buffer.capacity
	buffer.size--

	queueCopy = append([]Data(nil), buffer.buffer...)

	buffer.hasSpace.Signal()

	return nil
}
