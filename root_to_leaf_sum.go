/**
Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.

An example is the root-to-leaf path 1->2->3 which represents the number 123.

Find the total sum of all root-to-leaf numbers.

Note: A leaf is a node with no children.

Example:

Input: [1,2,3]
    1
   / \
  2   3
Output: 25
Explanation:
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Therefore, sum = 12 + 13 = 25.
**/


func dfs(node *TreeNode, val int, sum *int) {
    if node != nil {
        val += node.Val
    }
    
    if node.Left != nil {
        dfs(node.Left, val * 10, sum)
    }
    
    if node.Right != nil {
        dfs(node.Right, val * 10, sum)
    }
    
    if node.Left == nil && node.Right == nil {
        *sum += val
    }
}

func sumNumbers(root *TreeNode) int {
    sum := 0
    if root == nil {
        return sum
    }
    dfs (root, 0, &sum)
    
    return sum
}