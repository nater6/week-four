package piscine

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if data < root.Data {
		// If data is less then parent go left
		if root.Left != nil {
			root.Left.Parent = root
		}
		root.Left = BTreeInsertData(root.Left, data)
	} else {
		// If data is greater or equal to the parent then go right
		if root.Right != nil {
			root.Right.Parent = root
		}
		root.Right = BTreeInsertData(root.Right, data)
	}
	return root
}
