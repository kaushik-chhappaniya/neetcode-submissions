/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	countL := 1
	if root.Left != nil {
		countL += maxDepth(root.Left)
	}
	countR := 1
	if root.Right != nil {
		countR += maxDepth(root.Right)
	}
	return max(countL, countR)
}