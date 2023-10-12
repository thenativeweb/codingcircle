package playlist_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thenativeweb/codingcircle/playlist"
)

func TestRemoveKthLastElement(t *testing.T) {
	t.Run("remove last element", func(t *testing.T) {
		list :=
			playlist.NewNode(2,
				playlist.NewNode(3,
					playlist.NewNode(5,
						playlist.NewNode(7,
							playlist.NewNode(11, nil)))))

		updatedList := playlist.RemoveKthLastElement(list, 1)
		values := playlist.GetListValues(updatedList)

		assert.Equal(t, []int{2, 3, 5, 7}, values)
	})

	t.Run("remove first element", func(t *testing.T) {
		list :=
			playlist.NewNode(2,
				playlist.NewNode(3,
					playlist.NewNode(5,
						playlist.NewNode(7,
							playlist.NewNode(11, nil)))))

		updatedList := playlist.RemoveKthLastElement(list, 5)
		values := playlist.GetListValues(updatedList)

		assert.Equal(t, []int{3, 5, 7, 11}, values)
	})

	t.Run("remove k-th last element", func(t *testing.T) {
		list :=
			playlist.NewNode(2,
				playlist.NewNode(3,
					playlist.NewNode(5,
						playlist.NewNode(7,
							playlist.NewNode(11, nil)))))

		updatedList := playlist.RemoveKthLastElement(list, 2)
		values := playlist.GetListValues(updatedList)

		assert.Equal(t, []int{2, 3, 5, 11}, values)
	})

	t.Run("with too large k", func(t *testing.T) {
		list :=
			playlist.NewNode(2,
				playlist.NewNode(3,
					playlist.NewNode(5,
						playlist.NewNode(7,
							playlist.NewNode(11, nil)))))

		updatedList := playlist.RemoveKthLastElement(list, 17)

		assert.Nil(t, updatedList)
	})
}
