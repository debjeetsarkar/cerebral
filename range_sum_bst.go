/**
Given the root node of a binary search tree, return the sum of values of all nodes with value between L and R (inclusive).

The binary search tree is guaranteed to have unique values.

Example 1:
Input: root = [10,5,15,3,7,null,18], L = 7, R = 15
Output: 32

Example 2:
Input: root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
Output: 23


    10
    /\
  5    15
 /\    /\
3  7  n  18

**/


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

func sum (root *TreeNode, l, r int) int {
	current := root
	stk := &Stack{}
	sum := 0

	for current != nil || !stk.isEmpty() {
		for current != nil {
			stk.push(current)
			current = current.Left
		}
		popped := stk.pop()

		if popped.Val >= l && popped.Val <=r {
			sum += popped.Val
		}

		current = popped
		current = current.Right
	}

	return sum
}





