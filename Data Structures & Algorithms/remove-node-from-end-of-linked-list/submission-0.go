/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
	left := dummy
	right := head
	// We move the right ptr by n nodes
	for n >0 {
		right = right.Next
		n--
	}
	// move left ptr till right reaches nil- meaning right goes till end and left goes till end - n (i.e, len - n)
	for right != nil {
		left = left.Next
		right = right.Next
	}
	// so left is node we need to remove
	left.Next = left.Next.Next
	return dummy.Next
}
