/**
Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.

Example 1:

Input: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
Output: 1

**/


type Stack []*TreeNode

func(s *Stack) push (val *TreeNode) {
	(*s) = append(*s, val)
}

func(s *Stack) pop () *TreeNode{
	l := len(*s)
	popped := (*s)[l-1]
	(*s) = (*s)[:l-1]
	return popped
}

func(s *Stack) isEmpty () {
	return len(*s) < 1
}

func kthSmallest(root *TreeNode, k int) int {
	traversed := inOrderTraversal(root)
	l := len(traversed)
	if l < k-1 {
		return traversed[l]
	}

	return traversed[k-1]
}

func inOrderTraversal(root *TreeNode) []int {
	result := []int
	if root == nil {
		return result
	}

	current := root
	inorderStack := &Stack{}

	for current != nil || !inorderStack.isEmpty() {
		for current != nil {
			inorderStack.push(current)
			current = current.Left
		}
		current := inorderStack.pop()
		result = append(result, current)
		current = current.Right
	}

	return result
}