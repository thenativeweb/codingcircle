package busfactor_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/busfactor"
)

func TestGetMaximumLoad(t *testing.T) {
	t.Run("with a single event", func(t *testing.T) {
		now := time.Now()

		result := busfactor.GetMaximumLoad([]busfactor.Event{
			{Time: now, Type: "joined", Count: 3},
		})

		assert.Equal(t, 3, result.Count)
		assert.Equal(t, now, result.From)
		assert.Equal(t, now, result.Until)
	})

	t.Run("with multiple events", func(t *testing.T) {
		now := time.Now()

		result := busfactor.GetMaximumLoad([]busfactor.Event{
			{Time: now, Type: "joined", Count: 3},                      // 3
			{Time: now.Add(1 * time.Second), Type: "left", Count: 1},   // 2
			{Time: now.Add(2 * time.Second), Type: "joined", Count: 5}, // 7
			{Time: now.Add(3 * time.Second), Type: "left", Count: 4},   // 3
			{Time: now.Add(4 * time.Second), Type: "joined", Count: 2}, // 5
			{Time: now.Add(5 * time.Second), Type: "left", Count: 4},   // 1
		})

		assert.Equal(t, 7, result.Count)
		assert.Equal(t, now.Add(2*time.Second), result.From)
		assert.Equal(t, now.Add(3*time.Second), result.Until)
	})
}
