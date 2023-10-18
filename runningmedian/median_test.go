package runningmedian_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/runningmedian"
)

func TestGetMedians(t *testing.T) {
	t.Run("single element", func(t *testing.T) {
		sequence := []int{23}
		medians := runningmedian.GetMedians(sequence)

		assert.Equal(t, []float64{23}, medians)
	})

	t.Run("two elements", func(t *testing.T) {
		sequence := []int{23, 42}
		medians := runningmedian.GetMedians(sequence)

		assert.Equal(t, []float64{23, 32.5}, medians)
	})

	t.Run("multiple elements (odd)", func(t *testing.T) {
		sequence := []int{23, 42, 36, 35, 31, 27, 172}
		medians := runningmedian.GetMedians(sequence)

		assert.Equal(t, []float64{23, 32.5, 36, 35.5, 35, 33, 35}, medians)
	})

	t.Run("multiple elements (even)", func(t *testing.T) {
		sequence := []int{23, 42, 36, 35, 31, 27}
		medians := runningmedian.GetMedians(sequence)

		assert.Equal(t, []float64{23, 32.5, 36, 35.5, 35, 33}, medians)
	})
}
