/**
Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.


Input:

   1
 /   \
2     3
 \
  5

Output: ["1->2->5", "1->3"]


Approach

1. If root is not nil and has either left or right then set initial result string as root.Val + "->"
2. Push the root to the stack
3. While stack is not empty
 3.1 Pop from stack and perform these
 	3.1.1 If the popped has left or right present then add popped.Val + "->" to current_result, push to stack
 	3.1.2 If the popped has no children then add popped.Val to the current_result. Add current result to result array and flush current result.
4. Return result
**/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	result := []string{}

	if root == nil {
		return result
	}
	dfs(root, "", result)

	return result
}



func dfs(node *TreeNode, path string, result []string) {
	path += strconv.Itoa(node.Val)

	if node.Left == nil && node.Right == nil {
		return
	}

	if node.Left != nil {
		dfs(node.Left, path + "->", result)
	}

	if node.Right != nil {
		dfs(node.Right, path + "->", result)
	}
}




