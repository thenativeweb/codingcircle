package floyd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/floyd"
)

func TestHasCycle(t *testing.T) {
	t.Run("returns false for an empty list", func(t *testing.T) {
		assert.False(t, floyd.HasCycle(nil))
	})

	t.Run("returns false for a single node", func(t *testing.T) {
		list := floyd.NewList(2)

		assert.False(t, floyd.HasCycle(list))
	})

	t.Run("returns false for a list without a cycle", func(t *testing.T) {
		list := floyd.NewList(2, 3, 5, 7, 11)

		assert.False(t, floyd.HasCycle(list))
	})

	t.Run("returns true for a list with a cycle", func(t *testing.T) {
		list := floyd.NewList(2, 3, 5, 7, 11)

		last := list.Next.Next.Next.Next
		third := list.Next.Next
		last.Next = third

		assert.True(t, floyd.HasCycle(list))
	})
}
