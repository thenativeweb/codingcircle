package egyptianfractions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/egyptianfractions"
)

func TestReduce(t *testing.T) {
	fraction := egyptianfractions.Fraction{
		Numerator:   4,
		Denominator: 8,
	}

	reducedFraction := egyptianfractions.Reduce(fraction)

	assert.Equal(t, 1, reducedFraction.Numerator)
	assert.Equal(t, 2, reducedFraction.Denominator)
}
