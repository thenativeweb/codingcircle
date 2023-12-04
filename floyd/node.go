package floyd

type Node struct {
	Value int
	Next  *Node
}

func NewList(values ...int) *Node {
	if len(values) == 0 {
		return nil
	}

	head := &Node{values[0], nil}
	current := head

	for _, value := range values[1:] {
		current.Next = &Node{value, nil}
		current = current.Next
	}

	return head
}
