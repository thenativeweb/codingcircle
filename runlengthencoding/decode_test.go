package runlengthencoding_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/runlengthencoding"
)

func TestDecode(t *testing.T) {
	t.Run("decodes an empty input", func(t *testing.T) {
		result := runlengthencoding.Decode("")
		assert.Equal(t, "", result)
	})

	t.Run("decodes a single character", func(t *testing.T) {
		result := runlengthencoding.Decode("1A")
		assert.Equal(t, "A", result)
	})

	t.Run("decodes multiple characters", func(t *testing.T) {
		result := runlengthencoding.Decode("1A1B1C1D")
		assert.Equal(t, "ABCD", result)
	})

	t.Run("decodes repeated characters", func(t *testing.T) {
		result := runlengthencoding.Decode("1A2B3C4D")
		assert.Equal(t, "ABBCCCDDDD", result)
	})

	t.Run("decodes characters repeated more than 9 times", func(t *testing.T) {
		result := runlengthencoding.Decode("1A10B3C4D")
		assert.Equal(t, "ABBBBBBBBBBCCCDDDD", result)
	})

}
