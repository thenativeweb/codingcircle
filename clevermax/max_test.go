package clevermax_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/clevermax"
)

func TestMax(t *testing.T) {
	t.Run("returns the larger number as maximum", func(t *testing.T) {
		assert.Equal(t, 42, clevermax.Max(23, 42))
	})

	t.Run("returns the given number if both are equal", func(t *testing.T) {
		assert.Equal(t, 23, clevermax.Max(23, 23))
	})
}

func TestMaxBitShift(t *testing.T) {
	t.Run("returns the larger number as maximum", func(t *testing.T) {
		assert.Equal(t, 42, clevermax.MaxBitShift(23, 42))
	})

	t.Run("returns the given number if both are equal", func(t *testing.T) {
		assert.Equal(t, 23, clevermax.MaxBitShift(23, 23))
	})
}
