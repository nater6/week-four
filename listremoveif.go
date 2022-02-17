package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	if l.Head == nil {
		return
	}

	x := l.Head

	for x != nil {

		if x.Next.Data == data_ref {
			x.Next = x.Next.Next
		}
		x = x.Next

	}
}
