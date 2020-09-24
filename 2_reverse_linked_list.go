/**

Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL

I/P 1->2->3->4->5->NULL, ,m=2, n=4

1. 1->Null

2. 2->3->4->Null

3. 5->Null

4. Reverse 2 4->3->2->Null

5. Connect 1. with 4. 1->4->3->2

6. Connect 5 with 3. 1->4->3->2->5->Null

**/

type ListNode struct {
	Val int
	Next *ListNode
}


func reverse(head *ListNode) *ListNode{
	if head == nil || head.Next == nil {
		return head
	}

	var prev, current, upcoming *Listnode
	current = head
	upcoming = current.Next

	for current.Next != nil {
		current.Next = prev
		prev = current
		current = upcoming
		upcoming = upcoming.Next
	}

	current.Next = prev

	return current
}

func reverseBetween(head *ListNode, m,n int) *ListNode {
	if m == n || head == nil || head.Next == nil{
		return head
	}

	var  originalHead, prevM, prevN, lastPart, currentHead *ListNode

	prevM = head
	originalHead = head

	if m > 1 {
		i := 1
		for i < m && head.Next != nil {
			prevM = head
			head = head.Next
			i++
		}

		prevM.Next = nil
	}

	prevN = head
	lastPart = head

	j := m
	for j < n && prevN.Next != nil {
		prevN = prevN.Next
		lastPart = lastPart.Next
		j++
	}

	lastPart = lastPart.Next
	prevN.Next = nil

	currentHead = head

	if m == 1 {
		originalHead = reverse(head)
	} else {
		prevM.Next = reverse(head)	
	}

	currentHead.Next = lastPart
	

	return originalHead
}