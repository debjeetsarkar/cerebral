/**
Given an unsorted array of integers, find the length of longest increasing subsequence.

Example:

Input: [10,9,2,5,3,7,101,18]
Output: 4 
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4. 
Note:

There may be more than one LIS combination, it is only necessary for you to return the length.
Your algorithm should run in O(n2) complexity.
**/


func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lengthOfLIS(nums []int) int{
	l := len(nums)

	if l == 0  || l == 1{
		return l
	}

	dp := make([]int, l)
    longestIncreasingSubsequence := 0

	for i := 0; i < len(dp); i ++ {
		dp[i] = 1
	}
    
	for i := 1; i < len(dp); i ++ {
		for j := 0; j < i; j ++ {
			if nums[i] > nums[j] && dp[i] <= dp[j] {
				dp[i] = 1 + dp[j]
			}
		}

		longestIncreasingSubsequence = max(longestIncreasingSubsequence, dp[i])
	}

	return longestIncreasingSubsequence
}