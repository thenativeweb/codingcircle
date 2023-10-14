package runningmedian

import "container/heap"

type MinHeap []int

func NewMinHeap() *MinHeap {
	h := &MinHeap{}
	heap.Init(h)
	return h
}

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(value any) {
	*h = append(*h, value.(int))
}

func (h *MinHeap) Pop() any {
	oldHeap := *h
	valueIndex := len(oldHeap) - 1
	value := oldHeap[valueIndex]
	*h = oldHeap[:valueIndex]
	return value
}
