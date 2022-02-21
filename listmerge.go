package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head = l2.Head
		return
	}
	x := l1.Head

	for x.Next != nil {
		x = x.Next
	}
	x.Next = l2.Head
}
