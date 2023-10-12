package playlist

func GetListValues(head *Node) []int {
	node := head

	var values []int
	for node != nil {
		values = append(values, node.Value)
		node = node.Next
	}

	return values
}
