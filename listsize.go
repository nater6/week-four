package piscine

func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	}

	n := l.Head
	counter := 1

	for n.Next != nil {
		n = n.Next
		counter++
	}

	return counter
}
