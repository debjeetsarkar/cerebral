package main

import "fmt"

/**
Given a string containing only three types of characters: '(', ')' and '*', write a function to check whether this string is valid. We define the validity of a string by these rules:

Any left parenthesis '(' must have a corresponding right parenthesis ')'.
Any right parenthesis ')' must have a corresponding left parenthesis '('.
Left parenthesis '(' must go before the corresponding right parenthesis ')'.
'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.
An empty string is also valid.
Example 1:
Input: "()"
Output: True
Example 2:
Input: "(*)"
Output: True
Example 3:
Input: "(*))"
Output: True
 */


type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string{
	l := len(*s)
	if l ==0 {
		return ""
	}
	p := (*s)[l -1]
	*s = (*s)[:l -1]

	return p
}

func (s *Stack) Seek() string {
	if len(*s) == 0{
		return ""
	}

	return (*s)[len(*s) -1]
}

func validParanthesis(str string) bool {
	l := len(str)
	if l == 0 {
		return true
	}

	if l == 1 && str == "*"{
		return true
	}


return false
}

func main () {
	fmt.Println(validParanthesis("(*()"))
}