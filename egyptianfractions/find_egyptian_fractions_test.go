package egyptianfractions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/egyptianfractions"
)

func TestFindEgyptianFractions(t *testing.T) {
	t.Run("returns 1/2 and 1/6 for 2/3", func(t *testing.T) {
		fraction := egyptianfractions.Fraction{2, 3}

		result := egyptianfractions.FindEgyptianFractions(fraction)

		assert.Equal(
			t,
			[]egyptianfractions.Fraction{{1, 2}, {1, 6}},
			result,
		)
	})

	t.Run("returns 1/3, 1/11 and 1/231 for 3/7", func(t *testing.T) {
		fraction := egyptianfractions.Fraction{3, 7}

		result := egyptianfractions.FindEgyptianFractions(fraction)

		assert.Equal(
			t,
			[]egyptianfractions.Fraction{{1, 3}, {1, 11}, {1, 231}},
			result,
		)
	})
}
