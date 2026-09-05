package lamportsmutualexclusionalgorithm

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

type MessageHeap struct {
	mc *MessageCollection
}

func (mh *MessageHeap) Push(m Message) {
	heap.Push(mh.mc, m)
}

func (mh *MessageHeap) Pop() (Message, bool) {
	if mh.mc.Len() == 0 {
		return Message{}, false
	}

	return heap.Pop(mh.mc).(Message), true
}

// Unused functions, just implemented the wrappers anyway
func (mh *MessageHeap) Init() {
	heap.Init(mh.mc)
}

func (mh *MessageHeap) ModifyIndex(index int, msg Message) {
	if index < mh.mc.Len() {
		return
	}

	mh.mc.Set(index, msg)
	heap.Fix(mh.mc, index)
}

func (mh *MessageHeap) Remove(index int) (Message, bool) {
	if index >= mh.mc.Len() {
		return Message{}, false
	}

	return heap.Remove(mh.mc, index).(Message), true
}
