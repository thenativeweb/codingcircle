package runlengthencoding_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/runlengthencoding"
)

func TestEncode(t *testing.T) {
	t.Run("encodes an empty input", func(t *testing.T) {
		result := runlengthencoding.Encode("")
		assert.Equal(t, "", result)
	})

	t.Run("encodes a single character", func(t *testing.T) {
		result := runlengthencoding.Encode("A")
		assert.Equal(t, "1A", result)
	})

	t.Run("encodes multiple characters", func(t *testing.T) {
		result := runlengthencoding.Encode("ABCD")
		assert.Equal(t, "1A1B1C1D", result)
	})

	t.Run("encodes repeated characters", func(t *testing.T) {
		result := runlengthencoding.Encode("ABBCCCDDDD")
		assert.Equal(t, "1A2B3C4D", result)
	})

	t.Run("encodes characters repeated more than 9 times", func(t *testing.T) {
		result := runlengthencoding.Encode("ABBBBBBBBBBCCCDDDD")
		assert.Equal(t, "1A10B3C4D", result)
	})
}
