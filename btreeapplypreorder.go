package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		f(root.Data)
		BTreeApplyPostorder(root.Left, f)
		BTreeApplyPostorder(root.Right, f)
	}
}
