/**
206. Reverse Linked List

Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL

Approach:

NULL <- 1<- 2 <-3  <- 4  5->NULL
                      p  c  u

**/

package main 

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}


func reverseList(head *ListNode) *ListNode {

	if head == nil || head>next == nil {
		return head
	}

	var prev, current, upcoming ListNode
	current = head
	upcoming = head.Next

	for current.Next != nil {
		current.Next = prev
		prev = current
		current = upcoming
		upcoming = upcoming.Next
	}

	current.Next = prev

	return current
}
