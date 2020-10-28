/**
Given a string s, find the longest palindromic subsequence's length in s. You may assume that the maximum length of s is 1000.

Example 1:
Input:

"bbbab"
Output:
4
One possible longest palindromic subsequence is "bbbb".
 

Example 2:
Input:

"cbbd"
Output:
2
**/


func reverse(s string) string {
	reversed := ""
	for index, val := range s {
		reversed = val + reversed
	}

	return reversed
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}


func longestPalindromeSubseq(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	dp := make([][]int, l+1)
	for i := range dp {
		dp[i] = make([]int, l+1)
	}

	for i := 0; i < len(dp); i ++ {
		for j := 0; j < len(dp[0]); j ++ { 
			if i ==0 || j == 0 {
				dp[i][j] = 0
			}
		}
	}

	rs := reverse(s)

	for i := 1; i < len(dp); i ++ {
		for j := 1; j < len(dp[0]); j ++ {
			if s[i-1] == rs[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[l][l]
}
