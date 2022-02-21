package piscine

func ListMerge(l1 *List, l2 *List) {
	x := l1.Head

	for x.Next != nil {
		x = x.Next
	}
	x.Next = l2.Head
}
