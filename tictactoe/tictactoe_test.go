package tictactoe_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/tictactoe"
)

func TestHasWon(t *testing.T) {
	testCases := []struct {
		name   string
		fields []tictactoe.Field
		size   int
		hasWon bool
	}{
		{
			name:   "Returns false for an empty board.",
			fields: []tictactoe.Field{},
			size:   3,
			hasWon: false,
		},
		{
			name:   "Returns false for a board with a single field.",
			fields: []tictactoe.Field{{0, 0}},
			size:   3,
			hasWon: false,
		},
		{
			name:   "Returns true for a row on a 3x3 board.",
			fields: []tictactoe.Field{{0, 0}, {0, 1}, {0, 2}},
			size:   3,
			hasWon: true,
		},
		{
			name:   "Returns true for a column on a 3x3 board.",
			fields: []tictactoe.Field{{0, 0}, {1, 0}, {2, 0}},
			size:   3,
			hasWon: true,
		},
		{
			name:   "Returns true for the slash diagonal on a 3x3 board.",
			fields: []tictactoe.Field{{0, 1}, {0, 2}, {1, 0}, {1, 1}, {2, 0}, {2, 2}},
			size:   3,
			hasWon: true,
		},
		{
			name:   "Returns true for the backslash diagonal on a 3x3 board.",
			fields: []tictactoe.Field{{0, 0}, {0, 1}, {1, 1}, {1, 2}, {2, 0}, {2, 2}},
			size:   3,
			hasWon: true,
		},
		{
			name:   "Returns true for a row on a 4x4 board.",
			fields: []tictactoe.Field{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
			size:   4,
			hasWon: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			hasWon := tictactoe.HasWon(testCase.fields, testCase.size)
			assert.Equal(t, testCase.hasWon, hasWon)
		})
	}
}
