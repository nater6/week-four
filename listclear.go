package piscine

func ListClear(l *List) {
	f := l.Head
	for f.Next != nil {
		f.Next = nil
	}
	l.Head = nil
}