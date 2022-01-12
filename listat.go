package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	counter := 0
	if pos == 0 {
		return l
	} else {

		// Secnd Part

		f := l
		for counter != pos {
			f = f.Next
			counter++
		}
		return f
	}
}
