/**
Given two strings text1 and text2, return the length of their longest common subsequence.

A subsequence of a string is a new string generated from the original string with some characters(can be none) deleted without changing the relative order of the remaining characters. (eg, "ace" is a subsequence of "abcde" while "aec" is not). A common subsequence of two strings is a subsequence that is common to both strings.


If there is no common subsequence, return 0.

Input: text1 = "abcde", text2 = "ace" 
Output: 3  
Explanation: The longest common subsequence is "ace" and its length is 3.


**/

func max(a, b int) int{
	if a > b {
		return a
	}

	return b
}

func lcs(text1, text2 string, l1, l2 int, cache [][]int) int {
	if cache[l1][l2] != -1 {
		return cache[l1][l2]
	}

	if l1 == 0 || l2 == 0 {
		return 0
	}

	if text1[l1-1] == text2[l2-1] {
		cache[l1][l2] = 1 + lcs(text1, text2, l1-1, l2-1, cache)
		return cache[l1][l2]
	} else {
		cache[l1][l2] = max(lcs(text1, text2, l1, l2-1, cache), lcs(text1, text2, l1-1, l2, cache))		
		return cache[l1][l2]
	}
}


func longestCommonSubsequence(text1 string, text2 string) int {
    l1, l2 := len(text1), len(text2)
    cache := make([][]int, l1 + 1)

    for i:= range cache {
    	cache[i] = make([]int, l2 + 1)
    }

    for i:= 0 ; i < len(cache); i++ {
    	for j :=0 ; j < len(cache[0]); j++ {
    		cache[i][j] = -1
    	}
    }

    return lcs(text1, text2, l1, l2, cache)
}


