package autocomplete_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/autocomplete"
)

func getTrie() *autocomplete.Trie {
	var t = autocomplete.NewTrie()

	t.Add("the")
	t.Add("native")
	t.Add("web")
	t.Add("website")

	return t
}

func TestTrie(t *testing.T) {
	t.Run("prefix exists", func(t *testing.T) {
		words := getTrie()

		result := words.Search("t")
		sort.Strings(result)

		assert.Equal(t, []string{"the"}, result)
	})

	t.Run("prefix exists with multiple words", func(t *testing.T) {
		words := getTrie()

		result := words.Search("w")
		sort.Strings(result)

		assert.Equal(t, []string{"web", "website"}, result)
	})

	t.Run("prefix does not exist", func(t *testing.T) {
		words := getTrie()

		result := words.Search("x")
		sort.Strings(result)

		assert.Equal(t, []string{}, result)
	})

	t.Run("prefix is a word", func(t *testing.T) {
		words := getTrie()

		result := words.Search("native")
		sort.Strings(result)

		assert.Equal(t, []string{"native"}, result)
	})

	t.Run("prefix is empty", func(t *testing.T) {
		words := getTrie()

		result := words.Search("")
		sort.Strings(result)

		assert.Equal(t, []string{"native", "the", "web", "website"}, result)
	})
}
