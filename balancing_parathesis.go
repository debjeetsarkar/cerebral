package main

import (
	"fmt"
)

/*
* Create map of opening : closing braces
* For every item in the input string
	* Check if it is an opening brace
		* If yes, push it to a stack
		* else if closing brace
			* Pop element from stack
			* Check if the popped element is the opening brace of the incoming closing brace using the map
			* if yes, continue
			* else bail out.
		* At the end if the stack is empty the string is balanced

		* Edge cases
			* If the string is empty its balanced by default.
			* If the string has odd number of characters, its unbalanced.
 */


type Stk [] string

func(s *Stk) Push(val string) {
	*s = append(*s, val)
}

func(s *Stk) Pop() string{
	l := len(*s)
	if l == 0 { return "" }

	n := (*s)[l-1]
	*s = (*s)[: l -1]
	return n
}


func isValid(s string) bool {
	l := len(s)
	if l==0 { return true }
	if l%2 != 0 { return false }

	BRACES := map [string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	stack := &Stk{}
	for _, brace := range s {
		if _, ok := BRACES[string(brace)]; ok {
			stack.Push(string(brace))
		} else {
			popped := stack.Pop()
			if BRACES[popped] != string(brace) {
				return false
			}
		}
	}
	//
	if len(*stack) > 0 {
		return false
	}

	return true
}

func main() {
	fmt.Println(isValid("{{"))
}