package lamportsmutualexclusionalgorithm

/*
* This file contains a heap implementation using Go's container/heap interface
* The only structure intended for use is MessageHeap. Message Collection is an internal data structure supporting the heap.
* */

import (
	"container/heap"
)

type MessageCollection []Message

func (mc MessageCollection) Len() int {
	return len(mc)
}

func (mc MessageCollection) Less(i, j int) bool {
	return mc[i].CompareTo(mc[j]) < 0
}

func (mc MessageCollection) Swap(i, j int) {
	mc[i], mc[j] = mc[j], mc[i]
}

func (mc *MessageCollection) Push(x any) {
	*mc = append(*mc, x.(Message))
}

func (mc *MessageCollection) Pop() any {
	old := *mc
	n := len(old)
	x := old[n-1]
	*mc = old[0 : n-1]
	return x
}

func (mc *MessageCollection) Set(index int, msg Message) {
	if index >= mc.Len() {
		return
	}

	(*mc)[index] = msg
}

// MessageHeap is a heap implementation supporting messages
type MessageHeap struct {
	mc *MessageCollection
}

// Push pushes a message onto the heap
func (mh *MessageHeap) Push(m Message) {
	heap.Push(mh.mc, m)
}

// Pop removes the message at the top of the heap, and adjusts the heap. The second return is the status of the Pop operation, which will be false if the heap was empty.
func (mh *MessageHeap) Pop() (Message, bool) {
	if mh.mc.Len() == 0 {
		return Message{}, false
	}

	return heap.Pop(mh.mc).(Message), true
}

// Unused functions, just implemented the wrappers anyway

// Init initializes a heap (not intended for use as we assume that the heap property is always maintained)
func (mh *MessageHeap) Init() {
	heap.Init(mh.mc)
}

// ModifyIndex replaces the message at index with msg, and adjusts the heap
func (mh *MessageHeap) ModifyIndex(index int, msg Message) {
	if index < mh.mc.Len() {
		return
	}

	mh.mc.Set(index, msg)
	heap.Fix(mh.mc, index)
}

// Remove removes the message at index, and adjusts the heap
func (mh *MessageHeap) Remove(index int) (Message, bool) {
	if index >= mh.mc.Len() {
		return Message{}, false
	}

	return heap.Remove(mh.mc, index).(Message), true
}
