package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	x := l.Next

	for x != nil {

		if x.Data > x.Next.Data {
			x.Data, x.Next.Data = x.Next.Data, x.Data
		}
		x = x.Next
	}
	return l.Next
}
