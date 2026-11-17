/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
    res := []int{}
	// create a function for recursion
	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		// if node == nil return base condition
		if node == nil {
			return
		}
		// add to result arr
		res = append(res, node.Val)
		// call preorder func on the left then right node
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	return res
}
