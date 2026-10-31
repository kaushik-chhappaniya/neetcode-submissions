/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	arr := []int{}
	if head.Next == nil {
		return true
	}
	if head.Next.Next == nil {
	if	head.Val == head.Next.Val {
		return true 
		}
	}
	curr := head
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}
	l, r := 0, len(arr)-1
	for l < r {
		if arr[l] != arr[r] {
			return false
		}
		l++
		r--
	}
	return true
}
