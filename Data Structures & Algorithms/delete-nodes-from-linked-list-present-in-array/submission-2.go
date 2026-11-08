/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func modifiedList(nums []int, head *ListNode) *ListNode {
	removed := map[int]bool{}
	for _, value := range nums {
		removed[value] = true
	}
	dummy := &ListNode{Val: 0}
	tail, current := dummy, head
	for current != nil {
		following := current.Next
		current.Next = nil
		if !removed[current.Val] {
			tail.Next = current
			tail = current
		}
		current = following
	}
	return dummy.Next
}

// func modifiedList(nums []int, head *ListNode) *ListNode {
//     hm := make(map[int]bool) 
// 	// Create a hashmap of int and bool. store all values from nums as true
// 	for _, v := range nums {
// 		hm[v] = true
// 	}
// 	fmt.Println(hm)
// 	dummyHead := &ListNode{Val:0}
// 	tail, curr := dummyHead, head
// 	for curr != nil {
// 		// if node.val present in hashmap then remove it. else continue
// 		nx := curr.Next
// 		curr.Next = nil
// 		if !hm[curr.Val] {
// 			tail.Next = curr
// 			tail = curr
// 		}
// 		curr = nx
// 	}
// 	return dummyHead.Next
// }
