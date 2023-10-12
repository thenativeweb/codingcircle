package removekthlastelement_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/removekthlastelement"
)

func TestRemoveKthLastElement(t *testing.T) {
	t.Run("remove last element", func(t *testing.T) {
		list :=
			removekthlastelement.NewNode(2,
				removekthlastelement.NewNode(3,
					removekthlastelement.NewNode(5,
						removekthlastelement.NewNode(7,
							removekthlastelement.NewNode(11, nil)))))

		updatedList := removekthlastelement.RemoveKthLastElement(list, 1)
		values := removekthlastelement.GetListValues(updatedList)

		assert.Equal(t, []int{2, 3, 5, 7}, values)
	})

	t.Run("remove first element", func(t *testing.T) {
		list :=
			removekthlastelement.NewNode(2,
				removekthlastelement.NewNode(3,
					removekthlastelement.NewNode(5,
						removekthlastelement.NewNode(7,
							removekthlastelement.NewNode(11, nil)))))

		updatedList := removekthlastelement.RemoveKthLastElement(list, 5)
		values := removekthlastelement.GetListValues(updatedList)

		assert.Equal(t, []int{3, 5, 7, 11}, values)
	})

	t.Run("remove k-th last element", func(t *testing.T) {
		list :=
			removekthlastelement.NewNode(2,
				removekthlastelement.NewNode(3,
					removekthlastelement.NewNode(5,
						removekthlastelement.NewNode(7,
							removekthlastelement.NewNode(11, nil)))))

		updatedList := removekthlastelement.RemoveKthLastElement(list, 2)
		values := removekthlastelement.GetListValues(updatedList)

		assert.Equal(t, []int{2, 3, 5, 11}, values)
	})

	t.Run("with too large k", func(t *testing.T) {
		list :=
			removekthlastelement.NewNode(2,
				removekthlastelement.NewNode(3,
					removekthlastelement.NewNode(5,
						removekthlastelement.NewNode(7,
							removekthlastelement.NewNode(11, nil)))))

		updatedList := removekthlastelement.RemoveKthLastElement(list, 17)

		assert.Nil(t, updatedList)
	})
}
