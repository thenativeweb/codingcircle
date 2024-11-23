package parentheses_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/parentheses"
)

func TestIsBalanced(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		assert.True(t, parentheses.IsBalanced(""))
	})

	t.Run("balanced string", func(t *testing.T) {
		assert.True(t, parentheses.IsBalanced("()"))
		assert.True(t, parentheses.IsBalanced("(())"))
		assert.True(t, parentheses.IsBalanced("()()"))
		assert.True(t, parentheses.IsBalanced("(()())"))
		assert.True(t, parentheses.IsBalanced("((()))"))
	})

	t.Run("unbalanced string", func(t *testing.T) {
		assert.False(t, parentheses.IsBalanced("("))
		assert.False(t, parentheses.IsBalanced(")"))
		assert.False(t, parentheses.IsBalanced(")("))
		assert.False(t, parentheses.IsBalanced("())"))
		assert.False(t, parentheses.IsBalanced("(()"))
		assert.False(t, parentheses.IsBalanced("((())"))
	})

	t.Run("balanced string with wildcard", func(t *testing.T) {
		assert.True(t, parentheses.IsBalanced("*"))
		assert.True(t, parentheses.IsBalanced("(*)"))
		assert.True(t, parentheses.IsBalanced("(*))"))
		assert.True(t, parentheses.IsBalanced("(**))*"))
		assert.True(t, parentheses.IsBalanced("*((**))*"))
		assert.True(t, parentheses.IsBalanced("(***"))
	})

	t.Run("unbalanced string with wildcard", func(t *testing.T) {
		assert.False(t, parentheses.IsBalanced("*("))
		assert.False(t, parentheses.IsBalanced("*)("))
		assert.False(t, parentheses.IsBalanced("***("))
		assert.False(t, parentheses.IsBalanced("(***(()"))
		assert.False(t, parentheses.IsBalanced("(**("))
	})
}
