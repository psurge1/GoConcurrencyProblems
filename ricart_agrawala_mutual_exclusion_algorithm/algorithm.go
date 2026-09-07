// Package ricartagrawalamutualexclusionalgorithm contains an implementation of the ricart-agrawala mutual exclusion algorithm
package ricartagrawalamutualexclusionalgorithm

/*
* Ricart-Agrawala's mutual exclusion algorithm is a consensus algorithm for a distributed system, only allowing one system in a fully connected network to enter a critical section at a given time.
* This algorithm abides by the 3 properties of mutual exclusion
* - Safety: we never enter a bad state / race condition
* - Lifeliness: progress is eventually made in the system, no deadlocks occur
* - Fairness: some notion of fairness (ex: no process starves, resource allocation occurs on a rolling basis, etc.)
*
* This implementation represents processes as go routines. Messages between processes occur through channels.
* Thus, this is only a simulation of a distributed system using concurrency.
* Ricart-Agrwala's algorithm assumes the following
* - Complete fault tolerance
* - All messages reach their final destination
*
* The algorithm is as follows
* - Send RequestCS(cself, pself) to request permission from all other processes
* - On receiving RequestCS(ci, pi):
*	- if pself is in the critical section already, defer responses to received messages (in order of arrival) until we release CS
*	- if pself doesn't want the critical section, send GrantPermission(cself, pself) to pi
*	- if pself wants the critical section
*		- if (cself, pself) < (ci, pi), dont grant permission, defer response after pself enters and leaves the CS
*		- else, send GrantPermission(cself, pself) to pi
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
