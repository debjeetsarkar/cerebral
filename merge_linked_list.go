package main

/*
Merge two sorted linked lists and return it as a new sorted list.
The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/


/*
Input: 1->2->4, 1->3->4
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode{
	result :=  &ListNode{ Val : 0, Next: nil}
	ref := result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			result.Next = l1
			l1 = l1.Next
		} else {
			result.Next = l2
			l2 = l2.Next
		}
		result = result.Next
	}
	if l1 != nil {
		result.Next = l1
	}

	if l2 != nil{
		result.Next = l2
	}

	return ref.Next
}
