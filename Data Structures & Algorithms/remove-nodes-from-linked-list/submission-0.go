/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNodes(head *ListNode) *ListNode {
		current := head
		var previous *ListNode
		for current != nil {
			following := current.Next
			current.Next = previous
			previous = current
			current = following
		}
		var result *ListNode
		maximum := 0
		for previous != nil {
			following := previous.Next
			if previous.Val >= maximum {
				maximum = previous.Val
				previous.Next = result
				result = previous
			} else {
				previous.Next = nil
			}
			previous = following
		}
		return result
	}
