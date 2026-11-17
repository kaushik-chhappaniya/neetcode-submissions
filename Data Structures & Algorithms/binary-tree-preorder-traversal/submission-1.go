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
	st := []*TreeNode{} // create stack to add nodes
	curr := root // create ptr to traverse
	for curr != nil || len(st) > 0 { 
		if curr != nil { // if ptr not nil 
		// add current ptr value to arr -preorder root
			res = append(res, curr.Val)
			// push right node on stack
			st = append(st, curr.Right)
			// now we've right node on stack, move ptr to left
			curr = curr.Left
		} else {
			// if current is nil meaning we're at end
			// move curr to last of stack meaning the rightmost part
			curr = st[len(st)-1]
			// pop the last element from stack 
			st = st[:len(st)-1]
		}
	}
	return res
}
