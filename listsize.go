package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListSize(l *List) int {
	counter := 0
	if l.Head == nil {
		return 0
	} else {
		f := l.Head
		for f != nil {
			counter++
			f = f.Next
		}
		return counter
	}
}
