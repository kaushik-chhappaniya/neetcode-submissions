/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var inorder func(node *TreeNode)
	// create helper function
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left) // append left node
		res = append(res, node.Val) // append the parent node
		inorder(node.Right) // append right node
	}
	inorder(root)
return res
}
