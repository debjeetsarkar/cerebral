/**
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example:

Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]

Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.

**/

func min(x, y int) int {
    if x < y {
        return x
    }
    
    return y
}

func minPathSum(grid [][]int) int {
    l := len(grid)
    if l == 0  {
        return 0
    }
    
    rows := l
    columns := len(grid[0])
    
    dp := make([][]int, rows)
    
    for i := range grid {
        dp[i] = make([]int, columns)
    }
    
    for i := 0; i < rows; i++ {
        for j := 0; j < columns; j++ {
            if i == 0 && j == 0 {
                dp[i][j] = grid[i][j]
            } else if i == 0 {
                dp[i][j] = grid[i][j] + dp[i][j-1]
            } else if j == 0 {
                dp[i][j] = grid[i][j] + dp[i-1][j]
            } else {
                dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    
    return dp[rows-1][columns-1]
}