/**
Given an array of integers nums and a positive integer k, find whether it's possible to divide this array into k non-empty subsets whose sums are all equal.

 

Example 1:

Input: nums = [4, 3, 2, 3, 5, 2, 1], k = 4
Output: True
Explanation: It's possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal sums.
**/



func canPartitionKSubsets(nums []int, k int) bool {
    l := len(nums)

    if l == 0 {
    	return false
    }

    totalSum = 0 

    for i := 0; i < l; i ++ {
    	totalSum += nums[i]
    }

    if totalSum % k != 0 {
    	return false
    }

    targetSum := totalSum / k

    rows := l + 1
    columns := targetSum + 1

    dp := make([][]bool, rows)

    for i := range rows {
    	dp[i] = make([]bool, columns)
    }

    for i := 0; i < rows; i++ {
    	for j := 0 ; j < len(dp[0]); j ++ {
    		if j == 0 {
    			dp[i][j] = true
    		}
    	}
    }


    for i := 1; i < rows; i++ {
    	for j := 1 ; j < len(dp[0]); j ++ {
    		if nums[i-1] <= j {
    			dp[i][j] = dp[i-1][j] || dp[i-1][j- nums[i]]
    		} else {
    			dp[i][j] = dp[i-1][j]
    		}
    	}
    }

    return dp[l][targetSum]
}