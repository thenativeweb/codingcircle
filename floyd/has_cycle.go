package floyd

func HasCycle(head *Node) bool {
	if head == nil {
		return false
	}

	tortoise, hare := head, head
	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next

		if tortoise == hare {
			return true
		}
	}

	return false
}
