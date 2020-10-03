/**
Given a non-empty binary tree, find the maximum path sum.

For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.

Example 1:

Input: [1,2,3]

       1
      / \
     2   3

Output: 6
**/

import "math"


func max(a,b int) int {
	if a >= b {
		return a
	}

	return b
}

func dfs(node *TreeNode, answer *int) int{
	if node == nil {
		return 0
	}

	left := dfs(node.Left)
	right := dfs(node.Right)
	answer = max(answer, left + right + node.Val)
	return max(0, node.Val + max(left, right))
}


func maxPathSum(root *TreeNode) int {

	answer := math.MinInt32
	if root == nil {
		return answer
	}

	dfs(root, &answer)
	return answer
}