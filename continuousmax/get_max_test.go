package continuousmax_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/continuousmax"
)

func TestGetMax(t *testing.T) {
	values := []int{27, 9, 17, 2, 12, 8}
	windowSize := 3

	maxValues, err := continuousmax.GetMax(values, windowSize)

	assert.NoError(t, err)
	assert.Equal(t, []int{27, 17, 17, 12}, maxValues)
}
