package main

import "fmt"

/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:
Input: "a b c a c b b"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.


Input: "p w w k e w"
*/

func lengthOfLongestSubstring( str string) int{
	length := len(str)
	if length == 0 { return 0 }
	if length ==1 { return 1 }
	maxLength := 0
	start := 0
	slider:= 0
	visitedMap := make(map[string]int)
	for j,char := range str {
		foundIndex,ok := visitedMap[string(char)]
		if !ok {
			if maxLength < (j-start) + 1 {
				maxLength++
			}
		} else {
			if foundIndex >= start {
				start = foundIndex + 1
			} else {
				if maxLength < (j-start) + 1 {
					maxLength++
				}
			}
		}
		visitedMap[string(char)] = j
		slider = j
	}
	if maxLength < (slider-start) + 1 {
		maxLength++
	}
	return maxLength
}

func main () {
	fmt.Println(lengthOfLongestSubstring("abba"))
}