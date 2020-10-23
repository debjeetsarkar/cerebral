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

func longestCommonSubsequence(text1 string, text2 string) int {
    l1, l2 := len(text1), len(text2)

    if l1 == 0 || l2 == 0 {
    	return 0
    }

    dp := make([][]int, l1+1)

    for i := range dp {
    	dp[i] = make([]int, l2 +1)
    }

    for i := 0; i <len(dp); i ++ {
    	for j := 0; j < len(dp[0]); j++ {
    		if i == 0 || j ==0 {
    			dp[i][j] = 0 
    		}
    	}
    }


    for i := 1; i < len(dp); i ++ {
    	for j := 1; j < len(dp[0]); j ++ {
    		if text1[i-1] == text2[j-1] {
    			dp[i][j] = 1 + dp[i-1][j-1]
    		} else {
    			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
    		}
    	}
    }

    return dp[l1][l2]
}

