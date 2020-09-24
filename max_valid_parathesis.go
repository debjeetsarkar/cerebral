package main

import "fmt"

type stack []int

func (stk *stack) push(str int) {
	*stk = append(*stk, str)
}

func (stk *stack) pop() int {
	if len(*stk) == 0 {
		return -1
	}
	s := (*stk)[len(*stk) - 1]
	*stk = (*stk)[: len(*stk) - 1]
	return s
}

func (stk *stack) peek() int{
 if len(*stk) == 0 {
 	return -2
 }
  return (*stk)[len(*stk) - 1]
}

func longestValidParentheses(s string) int {
 maxLen := 0
 refStack := &stack{-1}

 if len(s) == 0 || len(s) == 1 {
 	return 0
 }

 for index, paranthesis := range s {
 	if string(paranthesis) == "(" {
 		refStack.push(index)
	} else if string(paranthesis) == ")" {
		refStack.pop()
		if refStack.peek() == -2 {
			refStack.push(index)
		} else {
			if maxLen < index - refStack.peek() {
				maxLen = index - refStack.peek()
			}
		}
	}
 }
 return maxLen
}



func main () {
	fmt.Println(longestValidParentheses("(()()"))
}