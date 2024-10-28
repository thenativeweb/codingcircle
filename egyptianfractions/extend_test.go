package egyptianfractions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/egyptianfractions"
)

func TestExtend(t *testing.T) {
	fractionA := egyptianfractions.Fraction{1, 3}
	fractionB := egyptianfractions.Fraction{4, 7}

	extendedFractionA, extendedFractionB :=
		egyptianfractions.Extend(fractionA, fractionB)

	assert.Equal(t, 7, extendedFractionA.Numerator)
	assert.Equal(t, 21, extendedFractionA.Denominator)

	assert.Equal(t, 12, extendedFractionB.Numerator)
	assert.Equal(t, 21, extendedFractionB.Denominator)
}
