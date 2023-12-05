package cons_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/cons"
)

func TestCar(t *testing.T) {
	car := cons.Car(cons.Cons(23, 42))

	assert.Equal(t, 23, car)
}

func TestCdr(t *testing.T) {
	car := cons.Cdr(cons.Cons(23, 42))

	assert.Equal(t, 42, car)
}
