/**
Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.

Example 1:
Input: "aba"
Output: True
Example 2:
Input: "abca"
Output: True
Explanation: You could delete the character 'c'.
**/

package main

import "fmt"


func isSubPartPalindrome(s string, left_index, right_index int) bool {
	for left_index < right_index {
		if string(s[left_index]) != string(s[right_index]) {
			return false
		} else {
			left_index ++
			right_index --
		}
	}

	return true
}

func validPalindrome(s string) bool {
	l := len(s)
	if l == 0 || l ==1 || l ==2 {
		return true
	}

	left, right := 0, l-1

	for left < right {
		if string(s[left]) == string(s[right]) {
			left ++
			right --
		} else {
			return isSubPartPalindrome(s, left, right -1) || isSubPartPalindrome(s, left + 1, right)
		}
	}

	return true
}


func main() {
	fmt.Println(validPalindrome("abbca"))
}



