package autocomplete

type Node struct {
	children    map[rune]*Node
	isEndOfWord bool
}

func NewNode() *Node {
	return &Node{
		children:    make(map[rune]*Node),
		isEndOfWord: false,
	}
}

func (n *Node) GetWords(prefix string) []string {
	var result []string

	if n.isEndOfWord {
		result = append(result, prefix)
	}

	for char, childNode := range n.children {
		result = append(
			result,
			childNode.GetWords(prefix+string(char))...,
		)
	}

	return result
}
