/**

Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

Note: A leaf is a node with no children.

Example:

Given the below binary tree and sum = 22,

      5
     / \
    4   8
   /   / \
  11  13  4
 /  \    / \
7    2  5   1

**/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



func dfs(node *TreeNode, sum int, result *[][]int, path []int) {
    if node != nil {
        path = append(path, node.Val)
    }
     
    if node.Left != nil {
        dfs(node.Left, sum-node.Val, result, path)
    }
    
    if node.Right != nil {
        dfs(node.Right, sum-node.Val, result, path)
    }
    
    if node.Left == nil && node.Right == nil && sum ==  node.Val {
        cp := make([]int, len(path))
        copy(cp, path)        
        *result = append(*result, cp)
    }
    
}

func pathSum(root *TreeNode, sum int) [][]int {
    result := [][]int{}
    
    if root == nil {
        return result
    }
    
    dfs(root, sum, &result, []int{})
    
    return result
}