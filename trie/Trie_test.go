package trie_test

import (
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/trie"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

var words = trie.New()

func setup() {
	words.Add("the")
	words.Add("native")
	words.Add("web")
	words.Add("website")
}

func TestTrie(t *testing.T) {
	t.Run("prefix exists", func(t *testing.T) {
		result := words.Search("t")
		sort.Strings(result)

		assert.Equal(t, []string{"the"}, result)
	})

	t.Run("prefix exists with multiple words", func(t *testing.T) {
		result := words.Search("w")
		sort.Strings(result)

		assert.Equal(t, []string{"web", "website"}, result)
	})

	t.Run("prefix does not exist", func(t *testing.T) {
		result := words.Search("x")
		sort.Strings(result)

		assert.Equal(t, []string{}, result)
	})

	t.Run("prefix is a word", func(t *testing.T) {
		result := words.Search("native")
		sort.Strings(result)

		assert.Equal(t, []string{"native"}, result)
	})

	t.Run("prefix is empty", func(t *testing.T) {
		result := words.Search("")
		sort.Strings(result)

		assert.Equal(t, []string{"native", "the", "web", "website"}, result)
	})
}
