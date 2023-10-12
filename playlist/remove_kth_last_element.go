package playlist

func RemoveKthLastElement(head *Node, k int) *Node {
	lead := head
	follow := head

	for i := 0; i < k; i++ {
		if lead == nil {
			return nil
		}
		lead = lead.Next
	}

	if lead == nil {
		return head.Next
	}

	for lead.Next != nil {
		lead = lead.Next
		follow = follow.Next
	}

	follow.Next = follow.Next.Next

	return head
}
