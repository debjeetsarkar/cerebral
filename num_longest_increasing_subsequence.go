/**
673. Number of Longest Increasing Subsequence

Add to List
Given an integer array nums, return the number of longest increasing subsequences.

Example 1:

Input: nums = [1,3,5,4,7]
Output: 2
Explanation: The two longest increasing subsequences are [1, 3, 4, 7] and [1, 3, 5, 7].
Example 2:

Input: nums = [2,2,2,2,2]
Output: 5
Explanation: The length of longest continuous increasing subsequence is 1, and there are 5 subsequences' length is 1, so output 5.
**/


func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func findNumberOfLIS(nums []int) int {
    l := len(nums)

    if l == 0 || l == 1 {
    	return l
    }

    dp := make([]int, l)
    count := make([]int, l)

    lis := 0
    ans := 0

    for i := 0; i < l; i ++ {
    	dp[i] = 1
    	count[i] = 1
    }

    for i := 1; i < l ; i ++ {
    	for j := 0; j < i ; j ++ {
    		if nums[i] > nums[j] {
    			if dp[i] <= dp[j] {
    				dp[i] = 1 + dp[j]
    				count[i] = count[j]
    			} else if dp[j] + 1 == dp[i] {
    				count[i] += count[j]
    			}
    		} 
    	}

        lis = max(lis, dp[i])
    }

    for i := 0; i < l; i ++ {
        if dp[i] == lis {
            ans += count[i]
        }
    }

    return  ans
}