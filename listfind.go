package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	x := l.Head

	for x != nil {
		if comp(x.Data, ref) {
			return &x.Data
		} else {
			x = x.Next
		}
	}

	return nil
}
