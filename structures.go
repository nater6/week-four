package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}
