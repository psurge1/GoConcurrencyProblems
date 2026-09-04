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
* */

var x = 0
