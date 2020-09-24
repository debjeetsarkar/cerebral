/**
In an alien language, surprisingly they also use english lowercase letters, but possibly in a different order. The order of the alphabet is some permutation of lowercase letters.

Given a sequence of words written in the alien language, and the order of the alphabet, return true if and only if the given words are sorted lexicographicaly in this alien language.

 

Example 1:

Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
Output: true
Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.
Example 2:

Input: words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
Output: false
Explanation: As 'd' comes after 'l' in this language, then words[0] > words[1], hence the sequence is unsorted.
Example 3:

Input: words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
Output: false
Explanation: The first three characters "app" match, and the second string is shorter (in size.) According to lexicographical rules "apple" > "app", because 'l' > '∅', where '∅' is defined as the blank character which is less than any other character


["apap","app"]
"abcdefghijklmnopqrstuvwxyz"

**/

package main

import (
	"fmt"
)


func compareWords(word1, word2 string, alienToHumanMap map[int]int) bool {
	l1 := len(word1)
	l2 := len(word2)

	i :=0
	j := 0

	difference_val := 0

	for i < l1 && j < l2 && difference_val == 0 {

		difference_val = alienToHumanMap[int(byte(word1[i]) - 'a')] - alienToHumanMap[int(byte(word2[j]) - 'a')]

		i++
		j++
	}

	if difference_val == 0 {
		if l1-l2 > 0 {
			return false	
		}
	} else if difference_val > 0 {
			return false
	}

	return true
}

func validateDict(words []string, order string) bool {
	l := len(words)

	if l == 0 {
		return false
	}
	if l == 1 {
		return true
	}

	alienToHumanMap := make(map[int]int)

	for index, value := range order {
		alienToHumanMap[int(byte(value) - 'a')] = index
	}

	for i := 1; i < len(words); i++ {
		if !compareWords(words[i-1], words[i], alienToHumanMap) {
			return false
		}
	}

	return true
}


func main () {
	fmt.Println(validateDict([]string{"aa","a"}, "abqwertyuioplkjhgfdszxcvnm"))
}
