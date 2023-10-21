package autocomplete

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

func (t *Trie) Add(word string) {
	currentNode := t.root

	for _, char := range word {
		if _, ok := currentNode.children[char]; !ok {
			currentNode.children[char] = NewNode()
		}
		currentNode = currentNode.children[char]
	}

	currentNode.isEndOfWord = true
}

func (t *Trie) Search(prefix string) []string {
	currentNode := t.root

	for _, char := range prefix {
		if _, ok := currentNode.children[char]; !ok {
			return []string{}
		}
		currentNode = currentNode.children[char]
	}

	return currentNode.GetWords(prefix)
}
