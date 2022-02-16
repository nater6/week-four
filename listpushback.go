package piscine

func ListPushBack(l *List, data interface{}) {
	adder := &NodeL{Data: data}

	// If l.head is nil make the first value data
	if l.Head == nil {
		l.Head = adder
		return
	}

	x := l.Head
	// Go through until the next value is nil
	for x.Next != nil {
		x = x.Next
	}

	x.Next = adder
}
