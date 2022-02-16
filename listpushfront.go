package piscine

func ListPushFront(l *List, data interface{}) {
	adder := NodeL{Data: data}

	if l.Head == nil {
		l.Head = &adder
	} else {
		adder.Next = l.Head
		l.Head = &adder
	}
}
