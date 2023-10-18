package runningmedian

import (
	"container/heap"
)

func getMedian(minHeap *MinHeap, maxHeap *MaxHeap) float64 {
	if minHeap.Len() > maxHeap.Len() {
		return float64((*minHeap)[0])
	}
	if minHeap.Len() < maxHeap.Len() {
		return float64((*maxHeap)[0])
	}
	return (float64((*minHeap)[0]) + float64((*maxHeap)[0])) / 2.0
}

func add(value int, minHeap *MinHeap, maxHeap *MaxHeap) {
	if minHeap.Len()+maxHeap.Len() < 1 {
		heap.Push(minHeap, value)
		return
	}

	median := getMedian(minHeap, maxHeap)
	if float64(value) > median {
		heap.Push(minHeap, value)
	} else {
		heap.Push(maxHeap, value)
	}
}

func rebalance(minHeap *MinHeap, maxHeap *MaxHeap) {
	if minHeap.Len() > maxHeap.Len()+1 {
		root := heap.Pop(minHeap).(int)
		heap.Push(maxHeap, root)
	} else if maxHeap.Len() > minHeap.Len()+1 {
		root := heap.Pop(maxHeap).(int)
		heap.Push(minHeap, root)
	}
}

func GetMedians(values []int) []float64 {
	minHeap := NewMinHeap()
	maxHeap := NewMaxHeap()

	medians := []float64{}

	for _, value := range values {
		add(value, minHeap, maxHeap)
		rebalance(minHeap, maxHeap)

		median := getMedian(minHeap, maxHeap)
		medians = append(medians, median)
	}

	return medians
}
