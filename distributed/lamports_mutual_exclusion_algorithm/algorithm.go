// Package lamportsmutualexclusionalgorithm contains an implementation of lamport's mutual exclusion algorithm
package lamportsmutualexclusionalgorithm

/*
* Lamport's mutual exclusion algorithm is a consensus algorithm for a distributed system, only allowing one system in a fully connected network to enter a critical section at a given time.
* This algorithm abides by the 3 properties of mutual exclusion
* - Safety: we never enter a bad state / race condition
* - Lifeliness: progress is eventually made in the system, no deadlocks occur
* - Fairness: some notion of fairness (ex: no process starves, resource allocation occurs on a rolling basis, etc.)
*
* This implementation represents processes as go routines. Messages between processes occur through channels.
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
*
* Testing (verified by error logging: processes pass logs to the log channel defined below)
* - Safety: only one process can enter the critical section
* - Liveliness: every process that requests the critical section eventually gets access to it
* - Fairness: processes enter the critical section in order or request time stamp (ordered by the happened before relation)
* */

import (
	"fmt"
	"sync"
)

func RunSimulation(N int) {
	wg := sync.WaitGroup{}
	logs := make(chan string, 1000)
	processes := InitSystem(N, logs)
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
		case msg := <-logs:
			fmt.Println(msg)
		case <-done:
			systemRunning = false
		default:
		}
	}

	close(logs)

	for log, ok := <-logs; ok; {
		fmt.Println(log)
		log, ok = <-logs
	}
}
