/**
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
**/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type Stack []*TreeNode

func (s *Stack) push(val *TreeNode) {
	*s = append(*s, val)
}

func (s *Stack) pop() (*TreeNode) {
	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return popped
}

func (s *Stack) isEmpty() bool {
	return len(*s) > 0
}


func inorderTraversal(root *TreeNode) []int {
	result := []int{}

	if root == nil {
		return result
	}

	inorderStack := &Stack{}
	current := root

	for current != nil || !inorderStack.isEmpty() {
		for current != nil {
			inorderStack.push(current)
			current = current.Left
		}
		current = inorderStack.pop()
		result = append(result, current.Val)
		current = current.Right
	}

	return result
}

func main () {
	inorderTraversal(nil)
}