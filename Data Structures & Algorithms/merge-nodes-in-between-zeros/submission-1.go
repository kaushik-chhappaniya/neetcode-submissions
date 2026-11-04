/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeNodes(head *ListNode) *ListNode {
    res := &ListNode{Val:0} // create result list
	h2 := res
	h1 := head.Next // start from next node as constraint says start,end always 0
	for h1 != nil && h1.Next != nil {
		if h1.Val == 0 && h1.Next != nil {
			// if h1.val == 0 then create new node and initialize with 0
			h2.Next = &ListNode{Val:0}
			h2 = h2.Next
		} else {
			// if not 0 then just add the value in result node's current value
			h2.Val += h1.Val
		}
		// keep iterating the h1
		h1 = h1.Next
	}
	return res
}
