/**
Given a string s of '(' , ')' and lowercase English characters. 

Your task is to remove the minimum number of parentheses ( '(' or ')', in any positions ) so that the resulting parentheses string is valid and return any valid string.

Formally, a parentheses string is valid if and only if:

It is the empty string, contains only lowercase characters, or
It can be written as AB (A concatenated with B), where A and B are valid strings, or
It can be written as (A), where A is a valid string.
 

Example 1:

Input: s = "lee(t(c)o)de)"
Output: "lee(t(c)o)de"
Explanation: "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.
Example 2:

Input: s = "a)b(c)d"
Output: "ab(c)d"
Example 3:

Input: s = "))(("
Output: ""
Explanation: An empty string is also valid.
Example 4:

Input: s = "(a(b(c)d)"
Output: "a(b(c)d)"

stack = []

**/

package main

import (
	"fmt"
)


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



func minRemoveToMakeValid(s string) string {
	l := len(s)
	if l ==0 {
		return ""
	}


	balance := 0
	refStack := &Stack{}

	for i :=0 ; i<l; i ++ {
		char := string(s[i])
		if char == "(" {
			balance++
			refStack.Push(char)
		} else if char == ")" && balance > 0 {
			balance--
			refStack.Push(char)
		} else if char != ")" && char != "("{
			refStack.Push(char)
		}
	}

	balance = 0

	len := len(*refStack)

	resultStr := ""

	fmt.Println(refStack)

	for len >0 {
		popped := refStack.Pop()

		if popped == ")" {
			balance--
			resultStr = popped + resultStr
		} else if popped == "(" && balance < 0{
			balance++
			resultStr = popped + resultStr
		} else if popped != ")" && popped != "("{
			resultStr = popped +resultStr
		}

		len--

	}

	return resultStr

}



func main() {
	fmt.Println(minRemoveToMakeValid("(a(b(c)d)"))
}

