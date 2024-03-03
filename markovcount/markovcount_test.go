package markovcount_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/markovcount"
)

func TestCount(t *testing.T) {
	t.Run("calculates the counts", func(t *testing.T) {
		transitions := []markovcount.Transition{
			{From: "a", To: "a", Probability: 0.9},
			{From: "a", To: "b", Probability: 0.075},
			{From: "a", To: "c", Probability: 0.025},
			{From: "b", To: "a", Probability: 0.15},
			{From: "b", To: "b", Probability: 0.8},
			{From: "b", To: "c", Probability: 0.05},
			{From: "c", To: "a", Probability: 0.25},
			{From: "c", To: "b", Probability: 0.25},
			{From: "c", To: "c", Probability: 0.5},
		}

		result := markovcount.Count("a", transitions, 10_000)

		assert.GreaterOrEqual(t, result["a"], 5_700)
		assert.LessOrEqual(t, result["a"], 6_800)

		assert.GreaterOrEqual(t, result["b"], 2_700)
		assert.LessOrEqual(t, result["b"], 3_600)

		assert.GreaterOrEqual(t, result["c"], 400)
		assert.LessOrEqual(t, result["c"], 900)
	})
}
