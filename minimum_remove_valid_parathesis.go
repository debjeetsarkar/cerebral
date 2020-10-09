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


type stack []string

func (s *stack) push(str string) {
    (*s) = append(*s, str)
}

func (s *stack) pop() string{
    if len(*s) == 0 {
        return ""
    }
    
    popped := (*s)[len(*s)-1]
    (*s) = (*s)[0:len(*s)-1]
    return popped
}

func minRemoveToMakeValid(s string) string {
    l := len(s)
    result := ""
    
    if l == 0 {
        return result
    }
    
    stk := &stack{}
    balance := 0
    
    for i := 0; i < l; i ++ {
        if string(s[i]) == "(" {
            stk.push(string(s[i]))
            balance++
        } else if string(s[i]) == ")" && balance > 0 {
            stk.push(string(s[i]))
            balance --
        } else if string(s[i]) != ")" && string(s[i]) != "("{
            stk.push(string(s[i]))
        }
    }
    
    balance = 0
    
    for l > 0 {
        popped := stk.pop()
        if popped == ")" {
            result = popped + result
            balance--
        } else if popped == "(" && balance < 0 {
            result = popped + result
            balance++
        } else if popped != "(" && popped != ")"{
            result = popped + result
        }
        
        l--;
    }
    
    return result
}


func main() {
	fmt.Println(minRemoveToMakeValid("(a(b(c)d)"))
}

