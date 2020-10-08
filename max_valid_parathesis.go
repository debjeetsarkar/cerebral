/**
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

 

Example 1:

Input: s = "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()".
**/

package main


import (
	"math"
	"fmt"
)

type stack []int

func (s *stack) push(a int) {
    *s = append(*s, a)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    
    return y
}

func (s *stack) pop() int {
    if len(*s) == 0 {
        return math.MinInt32
    }

    popped := (*s)[len(*s) - 1]
    (*s) = (*s)[0:len(*s) - 1]
    return popped
}

func (s *stack) peek() int {
     if len(*s) == 0 {
        return math.MinInt32
    }
    
    return (*s)[len(*s)-1]
}

func longestValidParentheses(s string) int {
    maxLength := 0
    l := len(s)
    if l == 0 {
       return  maxLength
    }
    
    stk := &stack{-1}
   
    for index, char := range s {
        if string(char) == "("{
            stk.push(index)
        } else if string(char) == ")" {
            stk.pop()
            if stk.peek() != math.MinInt32 {
                maxLength = max(maxLength, index - stk.peek())    
            } else {
                  stk.push(index)  
            }
        }
    }
    
    return  maxLength
}


func main () {
	fmt.Println(longestValidParentheses("(()()"))
}
