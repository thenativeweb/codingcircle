package heap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/heap"
)

func TestNew(t *testing.T) {
	t.Run("Creates an empty heap.", func(t *testing.T) {
		h := heap.New[int]()

		assert.True(t, h.IsValid())
		assert.Len(t, h.Items(), 0)
	})
}

func TestInsert(t *testing.T) {
	t.Run("Inserts a value into the heap.", func(t *testing.T) {
		h := heap.New[int]()

		h.Insert(23, 1)

		assert.True(t, h.IsValid())
		assert.Len(t, h.Items(), 1)
		assert.Equal(t, []heap.Item[int]{
			{Value: 23, Priority: 1},
		}, h.Items())
	})

	t.Run("Inserts multiple values into the heap.", func(t *testing.T) {
		h := heap.New[int]()

		h.Insert(42, 2)
		h.Insert(17, 3)
		h.Insert(23, 1)

		assert.True(t, h.IsValid())
		assert.Len(t, h.Items(), 3)
		assert.Equal(t, []heap.Item[int]{
			{Value: 23, Priority: 1},
			{Value: 17, Priority: 3},
			{Value: 42, Priority: 2},
		}, h.Items())
	})
}

func TestRemove(t *testing.T) {
	t.Run("Removes a value from the heap.", func(t *testing.T) {
		h := heap.New[int]()

		h.Insert(23, 1)

		err := h.Remove(0)
		assert.NoError(t, err)

		assert.True(t, h.IsValid())
		assert.Len(t, h.Items(), 0)
	})

	t.Run("Removes multiple values from the heap.", func(t *testing.T) {
		h := heap.New[int]()

		h.Insert(42, 2)
		h.Insert(17, 3)
		h.Insert(23, 1)

		err := h.Remove(1)
		assert.NoError(t, err)
		err = h.Remove(1)
		assert.NoError(t, err)

		assert.True(t, h.IsValid())
		assert.Len(t, h.Items(), 1)
		assert.Equal(t, []heap.Item[int]{
			{Value: 23, Priority: 1},
		}, h.Items())
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Updates a value in the heap.", func(t *testing.T) {
		h := heap.New[int]()

		h.Insert(42, 2)
		h.Insert(17, 3)
		h.Insert(23, 1)

		err := h.Update(2, 0)
		assert.NoError(t, err)

		assert.True(t, h.IsValid())
		assert.Len(t, h.Items(), 3)
		assert.Equal(t, []heap.Item[int]{
			{Value: 42, Priority: 0},
			{Value: 17, Priority: 3},
			{Value: 23, Priority: 1},
		}, h.Items())
	})
}

func TestPeek(t *testing.T) {
	t.Run("Returns the minimum value.", func(t *testing.T) {
		h := heap.New[int]()

		h.Insert(42, 2)
		h.Insert(17, 3)
		h.Insert(23, 1)

		assert.Equal(t, 23, h.Peek().Value)
		assert.Equal(t, 1, h.Peek().Priority)
	})

	t.Run("Returns nil if the heap is empty.", func(t *testing.T) {
		h := heap.New[int]()

		assert.Nil(t, h.Peek())
	})
}

func TestComplex(t *testing.T) {
	t.Run("Complex test.", func(t *testing.T) {
		h := heap.New[int]()

		h.Insert(27, 6)
		h.Insert(9, 3)
		h.Insert(17, 5)
		h.Insert(2, 1)
		h.Insert(12, 4)
		h.Insert(8, 2)

		assert.True(t, h.IsValid())
		assert.Len(t, h.Items(), 6)
		assert.Equal(t, []heap.Item[int]{
			{Value: 2, Priority: 1},
			{Value: 9, Priority: 3},
			{Value: 8, Priority: 2},
			{Value: 27, Priority: 6},
			{Value: 12, Priority: 4},
			{Value: 17, Priority: 5},
		}, h.Items())

		err := h.Remove(3)
		assert.NoError(t, err)

		assert.True(t, h.IsValid())
		assert.Len(t, h.Items(), 5)
		assert.Equal(t, []heap.Item[int]{
			{Value: 2, Priority: 1},
			{Value: 9, Priority: 3},
			{Value: 8, Priority: 2},
			{Value: 17, Priority: 5},
			{Value: 12, Priority: 4},
		}, h.Items())

		err = h.Update(3, -1)
		assert.NoError(t, err)

		assert.True(t, h.IsValid())
		assert.Len(t, h.Items(), 5)
		assert.Equal(t, []heap.Item[int]{
			{Value: 17, Priority: -1},
			{Value: 2, Priority: 1},
			{Value: 8, Priority: 2},
			{Value: 9, Priority: 3},
			{Value: 12, Priority: 4},
		}, h.Items())
	})
}
