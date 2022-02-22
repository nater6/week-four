package piscine

func ListReverse(l *List) {
	cur := l.Head
	var nex *NodeL
	var pre *NodeL

	for cur != nil {
		nex = cur.Next
		cur.Next = pre
		pre = cur

		if nex != nil {
			cur = nex
		} else {
			l.Head = cur
			// cur = nex
		}

	}
}
