package heap

import "errors"

type Item[TValue any] struct {
	Value    TValue
	Priority int
}

type MinHeap[TValue any] struct {
	items []Item[TValue]
}

func New[TValue any]() *MinHeap[TValue] {
	return &MinHeap[TValue]{}
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

func leftChildIndex(i int) int {
	return 2*i + 1
}

func rightChildIndex(i int) int {
	return 2*i + 2
}

func (h *MinHeap[TValue]) siftUp(index int) {
	for index > 0 {
		parent := parentIndex(index)

		if h.items[parent].Priority > h.items[index].Priority {
			h.items[parent], h.items[index] = h.items[index], h.items[parent]
			index = parent
		} else {
			break
		}
	}
}

func (h *MinHeap[TValue]) siftDown(index int) {
	left := leftChildIndex(index)
	right := rightChildIndex(index)
	smallest := index

	if left < len(h.items) && h.items[left].Priority < h.items[smallest].Priority {
		smallest = left
	}

	if right < len(h.items) && h.items[right].Priority < h.items[smallest].Priority {
		smallest = right
	}

	if smallest != index {
		h.items[index], h.items[smallest] = h.items[smallest], h.items[index]
		h.siftDown(smallest)
	}
}

func (h *MinHeap[TValue]) Peek() *Item[TValue] {
	if len(h.items) == 0 {
		return nil
	}

	return &h.items[0]
}

func (h *MinHeap[TValue]) IsValid() bool {
	for i := 1; i < len(h.items); i++ {
		if h.items[parentIndex(i)].Priority > h.items[i].Priority {
			return false
		}
	}
	return true
}

func (h *MinHeap[TValue]) Insert(value TValue, priority int) {
	h.items = append(h.items, Item[TValue]{Priority: priority, Value: value})
	h.siftUp(len(h.items) - 1)
}

func (h *MinHeap[TValue]) Remove(index int) error {
	if index < 0 || index >= len(h.items) {
		return errors.New("index out of range")
	}

	h.items[index] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]

	h.siftDown(index)

	return nil
}

func (h *MinHeap[TValue]) Update(index int, newPriority int) error {
	if index < 0 || index >= len(h.items) {
		return errors.New("index out of range")
	}

	oldPriority := h.items[index].Priority
	h.items[index].Priority = newPriority

	if newPriority < oldPriority {
		h.siftUp(index)
	} else {
		h.siftDown(index)
	}

	return nil
}

func (h *MinHeap[TValue]) Items() []Item[TValue] {
	return h.items
}
