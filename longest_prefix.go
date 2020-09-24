package main

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
*/

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}
	minLenWrd := strs[0]
	minLen := len(minLenWrd)
	// find minimum length word for prefix search
	for _, word := range strs {
		if len(word) < len(minLenWrd) {
			minLenWrd = word
			minLen = len(word)
		}
	}

	low := 1
	for low <= minLen {
		mid := (low + minLen) / 2
		if startsWith(minLenWrd, strs, mid) {
			low = mid + 1
		} else {
			minLen = mid - 1
		}
	}
	return minLenWrd[:low-1]
}

func startsWith (minLenWord string, strs []string, mid int) bool{
	substr := minLenWord[:mid]
	for _, word := range strs {
		if !strings.HasPrefix(word, substr) { return false }
	}
	return true
}

func main() {
	str := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(str))
}