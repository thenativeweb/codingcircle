package playlist

type Node struct {
	Value int
	Next  *Node
}

func NewNode(value int, next *Node) *Node {
	return &Node{
		Value: value,
		Next:  next,
	}
}
