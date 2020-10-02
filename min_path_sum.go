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


func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minPathSum(grid [][]int) int {
	l := len(grid)

	if l == 0 {
		return 0
	}

	h := l
	w := len(grid[0])

	dp := make([][]int, h)

	for i := range grid {
		dp[i] = make([]int, w)
	}

	for row := 0 ; row < h; row ++ {
		for column := 0; column < w; column ++ {
			if row == 0 && column == 0 {
				dp[row][column] = grid[row][column]
			} else if row == 0 {
				dp[row][column] = grid[row][column] + dp[row][column-1]
			} else if column == 0 {
				dp[row][column] = grid[row][column] + dp[row-1][column])
			} else {
				dp[row][column] = grid[row][column] + min(dp[row][column-1], dp[row-1][column])	
			}
		}	
	}

	return dp[h-1][w-1] 
}