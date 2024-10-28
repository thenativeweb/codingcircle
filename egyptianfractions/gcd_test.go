package egyptianfractions_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/egyptianfractions"
)

func TestGCD(t *testing.T) {
	t.Run("returns the greatest common divisor if one exists", func(t *testing.T) {
		result := egyptianfractions.GCD(8, 12)
		assert.Equal(t, 4, result)
	})

	t.Run("returns 1 if no greatest common divisor exists", func(t *testing.T) {
		result := egyptianfractions.GCD(2, 3)
		assert.Equal(t, 1, result)
	})
}
