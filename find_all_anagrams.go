/**

Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

The order of output does not matter.

Example 1:

Input:
s: "cbaebabacd" p: "abc"

Output:
[0, 6]

Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
**/

package main

import (
	"fmt"
)

func findAnagrams(s, p string) []int {
	l := len(s)

	start := 0
	window := make(map[string]int)
	counter := 0

	resultArr := []int{}

	if l < len(p) {
		return resultArr
	}

	// populate the window map
	for _, val := range p {
		v, ok := window[string(val)]
		if ok {
			window[string(val)] = v + 1
		} else {
			window[string(val)] = 1	
			counter++
		}
	}
	
	for slider := 0; slider < l; slider++ {
		char := string(s[slider])
		val, ok := window[char]
		if ok {
			window[char] = val - 1
			if window[char] == 0 {
				counter--
			}
		}

		fmt.Println("slider", slider, counter)

		// found all the characters required, now check for anagram presence
		for counter == 0 {	
			if (slider - start + 1)  == len(p) {
				resultArr = append(resultArr, start)
			} 

			startChar := string(s[start])
			val, ok := window[startChar]

			if ok {
				window[startChar] = val + 1
				if window[startChar] > 0 {
					counter++
				}
			}

			start++
		} 
	}

	return resultArr
}

/**
"baa"
"aa"
**/

func main () {
	fmt.Println(findAnagrams("aabab", "ab"))
}