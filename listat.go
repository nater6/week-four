package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	counter := 0
	f := l
	for f != nil {
		if counter == pos {
			return f
		}
		counter++
		f = f.Next
	}
	return nil
}
