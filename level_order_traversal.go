/**
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7]
**/

type Queue []*TreeNode

func (q *Queue) push(node *TreeNode) {
	*q = append(*q, node)
}

func (q *Queue) pop() *TreeNode{
	popped := (*q)[0]
	*q = (*q)[1: len(*q)]
	return popped
}

func(q *Queue) isEmpty() bool {
	return len(*q) < 1
}


func levelOrder (root *TreeNode) [][]int {
	result := [][]int {}
	if root == nil {
		return result
	}

	q := &Queue{root}

	for !q.isEmpty() {
		l := len(*q)

		local_result := []int {}
		for i := 0 ; i < l; i++ {
			popped := q.Pop()
			local_result = append(local_result, popped.Val)
			if popped.Left != nil {
				q.Push(popped.Left)
			}

			if popped.Right != nil {
				q.Push(popped.Right)
			}
		}
		result = append(result, local_result)
	}

	return result
}