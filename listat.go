package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	counter := 0
	f := l
	for counter != pos {
		f = f.Next
		counter++
	}
	if counter == pos {
		return f
	}

	return nil
}
