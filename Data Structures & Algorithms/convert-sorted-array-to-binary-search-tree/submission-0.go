func sortedArrayToBST(nums []int) *TreeNode {
	var helper func(l,r int) *TreeNode 
	// create helper function
	helper = func(l, r int) *TreeNode {
		// base case where left > right return nil
		if l > r {
			return nil
		}
		// find middle index as the array is sorted
		m := (l + r) / 2
		// create a parent root node with the middle index
		root := &TreeNode{Val: nums[m]}
		// call helper on the same with left and middle(current)-1)
		root.Left = helper(l, m-1)
		// call helper on the same with middle(current)+1 and right)
		root.Right = helper(m+1, r)
		return root
	}
	return helper(0, len(nums)-1)
}
