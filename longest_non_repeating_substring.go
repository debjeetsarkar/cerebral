package main

import "fmt"

/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:
Input: "abcadfg"
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


Input: "p w w k e w p"
*/

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func lengthOfLongestSubstring(s string) int {
    l := len(s)
    visitedMap := make(map[string]int)
    maxLength := 0
    start := 0
    for index, char := range s {
        foundIndex, ok := visitedMap[string(char)]
        if !ok {
            maxLength = max(maxLength, index - start + 1)
        } else {
            if foundIndex >= start {
                start = foundIndex + 1
            } else {
                maxLength = max(maxLength, index - start + 1)
            }
        }
        
        visitedMap[string(char)] = index
    }
    
    maxLength = max(maxLength, l - start)
    
    return maxLength
}
func main () {
	fmt.Println(lengthOfLongestSubstring("abcadfg"))
}