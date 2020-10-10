/**
Given an integer array with all positive numbers and no duplicates, find the number of possible combinations that add up to a positive integer target.

Example:

nums = [1, 2, 3]
target = 4

The possible combination ways are:
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)

Note that different sequences are counted as different combinations.

Therefore the output is 7.
**/


func combinationSum4(nums []int, target int) int {
    if len(nums) == 0 {
        return 0
    }
    dp := make([]int, target + 1)
    dp[0] = 1
    for i := 1; i <=target; i ++ {
        for _, num := range nums {
            if i < num {
                continue
            }
            dp[i] = dp[i] + dp[i-num]    
        }
    }
    
    return dp[target]
}