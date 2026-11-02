/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil{
		return head
	}
	for head != nil && head.Val == val {
        head = head.Next
    }
	prev, ptr := head, head
	for ptr != nil {
		   if ptr.Val == val {
            prev.Next = ptr.Next
            ptr = ptr.Next
        } else {
            prev = ptr
            ptr = ptr.Next
        }
	}
	return head
}
