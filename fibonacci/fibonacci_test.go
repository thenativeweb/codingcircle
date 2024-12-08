package fibonacci_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/fibonacci"
)

func TestFib(t *testing.T) {
	assert.Equal(t, 1, fibonacci.Fib(1))
	assert.Equal(t, 1, fibonacci.Fib(2))
	assert.Equal(t, 2, fibonacci.Fib(3))
	assert.Equal(t, 3, fibonacci.Fib(4))
	assert.Equal(t, 5, fibonacci.Fib(5))
	assert.Equal(t, 8, fibonacci.Fib(6))
	assert.Equal(t, 13, fibonacci.Fib(7))
	assert.Equal(t, 21, fibonacci.Fib(8))
	assert.Equal(t, 34, fibonacci.Fib(9))
	assert.Equal(t, 55, fibonacci.Fib(10))

	assert.Equal(t, 1134903170, fibonacci.Fib(45))
}
