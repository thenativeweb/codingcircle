package trie_test

import (
	"sort"
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
		name   string
		prefix string
		want   []string
	}{
		{
			name:   "prefix exists",
			prefix: "t",
			want:   []string{"the"},
		},
		{
			name:   "prefix exists with multiple words",
			prefix: "w",
			want:   []string{"web", "website"},
		},
		{
			name:   "prefix does not exist",
			prefix: "x",
			want:   []string{},
		},
		{
			name:   "prefix is a word",
			prefix: "native",
			want:   []string{"native"},
		},
		{
			name:   "prefix is empty",
			prefix: "",
			want:   []string{"native", "the", "web", "website"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := words.Search(test.prefix)
			sort.Strings(result)

			assert.Equal(t, test.want, result)
		})
	}
}
