/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    	list1_head := list1
	for i := 0; i < a-1; i++ {
		list1 = list1.Next
	}
	list1_ptA := list1
	fmt.Println("list1_ptA.Val  ", list1_ptA.Val)
	for i := a; i < b+1 && list1 != nil; i++ {
		list1 = list1.Next
	}
	list1_ptB := list1.Next
	list1_ptA.Next = list2
	fmt.Println("list1_ptB.Val  ", list1_ptB.Val)
	// fmt.Println("list1_ptA.NextVal  ",list1_ptA.Next.Val)
	for list1_ptA.Next != nil {
		list1_ptA = list1_ptA.Next
	}
	list1_ptA.Next = list1_ptB

	return list1_head
	}