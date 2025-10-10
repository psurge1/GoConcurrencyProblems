package producer_consumer_problem

import (
	"fmt"
	"math/rand"
	"sync"
)

func Test() {
	bufferWidth := 4
	numProducers := 16
	numConsumers := 16

	fmt.Printf("Producer Consumer Problem: (%d producers) (%d consumers) (%d spots in buffer)\n", numProducers, numConsumers, bufferWidth)

	if numProducers <= numConsumers+bufferWidth &&
		bufferWidth > 0 {
		testCase := TestParams(bufferWidth, numProducers, numConsumers)
		testCase()
	} else {
		fmt.Printf("Starvation! Buffer too small and/or too many producers! Cannot run test!\n")
	}
}

func TestParams(bufferWidth int, numProducers int, numConsumers int) func() {
	buffer := CreatePCQueue(bufferWidth)
	wg := sync.WaitGroup{}

	var threads []func() error
	for range numProducers {
		threads = append(threads, buffer.produce)
	}
	for range numConsumers {
		threads = append(threads, buffer.consume)
	}
	// fmt.Println("Threads:", threads)

	rand.Shuffle(len(threads), func(i, j int) {
		threads[i], threads[j] = threads[j], threads[i]
	})
	// fmt.Println("Shuffled Threads:", threads)

	return func() {
		for i, function := range threads {
			wg.Go(func() {
				if function() != nil {
					fmt.Printf("Thread %d caused an error!\n", i)
				}
			})
		}
		wg.Wait()
	}
}
