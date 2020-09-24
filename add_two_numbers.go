package main
/**
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	result := &ListNode{}
	head := result

	for l1 != nil || l2 != nil {
		val1 := 0
		val2 := 0
		if l1 != nil && l1.Val != 0 {
			val1 = l1.Val
		}

		if l2 != nil && l2.Val != 0 {
			val2 = l2.Val
		}

		sum := carry + val1 + val2

		node := &ListNode{}
		if sum > 9 {
			carry = sum /10
			node.Val = sum % 10
		} else {
			node.Val = sum
			carry = 0
		}

		result.Next = node
		result = result.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

	}

	if carry > 0 {
		node := &ListNode{}
		node.Val = carry
		result.Next = node
	}

	return head.Next
}