package trie_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/trie"
)

func TestTrie(t *testing.T) {
	words := trie.New()

	words.Add("the")
	words.Add("native")
	words.Add("web")
	words.Add("website")

	tests := []struct {
		name     string
		prefix   string
		expected []string
	}{
		{
			name:     "prefix exists",
			prefix:   "t",
			expected: []string{"the"},
		},
		{
			name:     "prefix exists with multiple words",
			prefix:   "w",
			expected: []string{"web", "website"},
		},
		{
			name:     "prefix does not exist",
			prefix:   "x",
			expected: []string{},
		},
		{
			name:     "prefix is a word",
			prefix:   "native",
			expected: []string{"native"},
		},
		{
			name:     "prefix is empty",
			prefix:   "",
			expected: []string{"the", "native", "web", "website"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := words.Search(test.prefix)
			assert.Equal(t, test.expected, actual)
		})
	}
}
