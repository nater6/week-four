package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	if l.Head == nil {
		return
	}

	curr := l.Head
	var prev *NodeL

	for curr != nil {
		if curr.Data != data_ref {
			prev = curr
			curr = curr.Next
		} else {
			prev.Next = curr.Next
			curr = curr.Next
		}
	}
}
