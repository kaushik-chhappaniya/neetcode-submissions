/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
		return head
	}
	curr := head
	for curr != nil && curr.Next != nil {
		// if curr val is same as nextnode val then just remap the next to next-next
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
 		}  else {
			curr = curr.Next
		}
	}
	return head
}
