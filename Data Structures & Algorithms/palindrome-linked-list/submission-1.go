/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head

	// find middle spot of the list --slow pointer
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse second half of the list
	var prev *ListNode 
	for slow != nil {
		tmp := slow.Next // temp node for reversing
		slow.Next = prev
		prev = slow
		slow = tmp
	}

	// check palindrome from 1st list and 2nd list
	l, r := head, prev
	for r != nil {
		if l.Val != r.Val {
			return false
		}
		l = l.Next
		r = r.Next
	}
	return true
}
