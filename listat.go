package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos == 0 {
		return l
	}

	x := l.Next
	y := 1

	if x == nil {
		return nil
	}

	for x != nil {
		if y == pos {
			break
		}
		x = x.Next
		y++

	}

	return x
}
