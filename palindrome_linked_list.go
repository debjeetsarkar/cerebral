/**

Given a singly linked list, determine if it is a palindrome.

Example 1:

Input: 1->2
Output: false
Example 2:

Input: 1->2->2->1
Output: true
**/


package main

import (
	"fmt"
)

type ListNode struct {
	Value int
	Next *ListNode
}

type Stack []int

func (s *Stack) push(value int) {
	*s = append(*s, value)
}

func (s *Stack) pop () int {
	l := len(*s)
	if l == 0 {
		return -10
	}

	popped := (*s)[l-1]
	*s = (*s)[:l-1]

	return popped
}


func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	stk := &Stack{}

	dummy := head
	for head != nil {
		stk.push(head.Value)
		head = head.Next
	}

	fmt.Println(dummy.Value)

	for dummy != nil {
		if dummy.Value != stk.pop() {
			return false
		}

		dummy = dummy.Next
	}

	return true
}



func main () {


    arr := []int{2}

    ll := &ListNode{Value: 1, Next: nil,}

    head := ll

    for _,val := range arr {
    	newNode := &ListNode{
    		Value: val,
    		Next: nil,
    	}

    	ll.Next = newNode
    	ll = ll.Next
    }

	fmt.Println(isPalindrome(head))
}