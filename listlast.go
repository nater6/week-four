package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	} else {
		f := l.Head
		for f.Next != nil {
			f = f.Next
		}
		return f.Data
	}
}
