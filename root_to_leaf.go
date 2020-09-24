/**
Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.

Example:

Input:

   1
 /   \
2     3
 \
  5
**/



type stack []int

func (stk *stack) push(str int) {
	*stk = append(*stk, str)
}

func (stk *stack) pop() int {
	s := (*stk)[len(*stk) - 1]
	*stk = (*stk)[0 : len(*stk) - 1]
	return s
}


func binaryTreePaths(root *TreeNode) []string {
	result := []string
	if root == nil {
		return result
	}


}

