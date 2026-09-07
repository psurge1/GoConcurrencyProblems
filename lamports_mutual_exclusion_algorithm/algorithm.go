// Package lamportsmutualexclusionalgorithm contains an implementation of lamport's mutual exclusion algorithm
package lamportsmutualexclusionalgorithm

/*
* Lamport's mutual exclusion algorithm is a consensus algorithm for a distributed system, only allowing one system in a fully connected network to enter a critical section at a given time.
* This algorithm abides by the 3 properties of mutual exclusion
* - Safety: we never enter a bad state / race condition
* - Lifeliness: progress is eventually made in the system, no deadlocks occur
* - Fairness: some notion of fairness (ex: no process starves, resource allocation occurs on a rolling basis, etc.)
*
* This implementation represents processes as go routines. Messages between processes occur througmc channels.
* Thus, this is only a simulation of a distributed system using concurrency.
* However, its important to keep in mind that Lamport's mutual exclusion algorithm is theoretical in nature,
* since it assumes FIFO message ordering,
* - reliable messages,
* - no faulty processes,
* - no malicious processes,
* - etc.
* All of which are not representative of real world constraints.
*
* The algorithm is as follows
* - Whenever a process wants to enter the critical section, it sends a request to all other messages, and adds its request to its internal queue
* - Whenever it recieves an acknowledgement from all other processes AND is at the front of the queue, it may enter the critical section
* Implementation
* - On sending ReqeustCS(cself, pself), add (pself, cself) to priority queue
* - On recieving RequestCS(ci, pi), add (ci, pi) to our priority queue, respond to pi with acknowledgement
* - On recieving ReleaseCS(ci, pi), remove pi from our priority queue
* */

import (
	"fmt"
	"sync"
)

func RunSimulation(N int) {
	wg := sync.WaitGroup{}
	messages := make(chan string, 1000)
	processes := InitSystem(N, messages)
	done := make(chan struct{})
	for _, proc := range processes {
		wg.Go(proc.Run)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	systemRunning := true
	for systemRunning {
		select {
		case msg := <-messages:
			fmt.Println(msg)
		case <-done:
			systemRunning = false
		default:
		}
	}

	close(messages)

	for msg, ok := <-messages; ok; {
		fmt.Println(msg)
		msg, ok = <-messages
	}
}
