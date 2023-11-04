package radixsort_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/radixsort"
)

func TestSort(t *testing.T) {
	unsorted := []int{12, 27, 2, 17, 9, 8}
	sorted := []int{2, 8, 9, 12, 17, 27}

	assert.Equal(t, sorted, radixsort.Sort(unsorted))
}
