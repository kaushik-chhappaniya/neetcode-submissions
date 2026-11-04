/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeNodes(head *ListNode) *ListNode {
    res := &ListNode{Val:0}
	h2 := res
	h1 := head.Next
	for h1 != nil && h1.Next != nil {
		if h1.Val == 0 && h1.Next != nil {
			h2.Next = &ListNode{Val:0}
			h2 = h2.Next
		} else {
			h2.Val += h1.Val
		}
		h1 = h1.Next
	}
	return res
}
