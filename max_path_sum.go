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

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func maxSumGoingDown(node *TreeNode, ans *int) int{
    if node == nil {
        return 0
    }
    
    left := maxSumGoingDown(node.Left, ans)
    right := maxSumGoingDown(node.Right, ans)
    
    *ans  = max(*ans, node.Val + left + right)
    
    return max(0, node.Val + max(left,right))
}

func maxPathSum(root *TreeNode) int {
    ans := math.MinInt32
    if root == nil {
        return 0
    }
    
    maxSumGoingDown(root, &ans)
    return ans
}