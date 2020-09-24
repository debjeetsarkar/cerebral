/**
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Example 2:

Input: "race a car"
Output: false
**/



package main

import (
	"fmt"
	"strings"
)


func isLetter(r rune) bool {
  return ((r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z'))
}

func checkPalindrome(s string) bool {
 	l := len(s)
 	if l == 0 {
 		return true
 	}

 	left := 0
 	right := l-1

	for left < right {

		if !isLetter(rune(s[left])) {
			left ++
			continue
		}

		if !isLetter(rune(s[right])) {
			right --
			continue
		}


		if isLetter(rune(s[left])) && isLetter(rune(s[right])){
			if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
				return false
			}

			left++
			right--
		} 
	}

	return true
}


func main () {
	fmt.Println(checkPalindrome( "race a car"))
}


