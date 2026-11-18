/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    // Create arr for result
	res := []int{}
	// create func for recursion
	var postOrder func(node *TreeNode) 
	postOrder = func(node *TreeNode) {
		if node == nil {
			return 
		}
		postOrder(node.Left)
		postOrder(node.Right)
		res = append(res, node.Val)
	}
	postOrder(root)
	return res
}
