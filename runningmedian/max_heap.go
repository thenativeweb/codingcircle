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
	valueIndex := len(oldHeap) - 1
	value := oldHeap[valueIndex]
	*h = oldHeap[:valueIndex]
	return value
}
