package runningmedian

import "container/heap"

type MaxHeap []int

func NewMaxHeap() *MaxHeap {
	h := &MaxHeap{}
	heap.Init(h)
	return h
}

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(value any) {
	*h = append(*h, value.(int))
}

func (h *MaxHeap) Pop() any {
	oldHeap := *h

	n := len(oldHeap)
	value := oldHeap[n-1]
	*h = oldHeap[0 : n-1]

	return value
}
