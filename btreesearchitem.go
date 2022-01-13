package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {

	if root == nil || root.Data == elem {
		return root
	} else {
		if elem < root.Data {
			// go left
			 return BTreeSearchItem(root.Left, elem)
		} else {
			// Go right
			 return BTreeSearchItem(root.Right, elem)
		}
		
	}
}