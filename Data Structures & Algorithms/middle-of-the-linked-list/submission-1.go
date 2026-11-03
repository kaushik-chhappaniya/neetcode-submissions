/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast !=nil && fast.Next != nil {
		// if u use fast.next.next it'll just 3 nodes, curr, next, next
		slow = slow.Next
		fast = fast.Next.Next
	}	
	return slow
}
