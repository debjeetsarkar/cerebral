/**
1302. Deepest Leaves Sum

Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
Output: 15
**/


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Queue []int

func (q *Queue) Push(val int) {
	*q = append(*q, val)
}

func (q *Queue) Pop() int {
	popped := (*q)[0]
	(*q) = (*q)[1: len(*q)]
	return popped
}

func (q *Queue) isEmpty() bool {
	return len(*q) < 1
}

func deepestLeavesSum(root *TreeNode) int {
	sum := 0
	if root == nil {
		return sum
	}

	// Perform BFS on tree
	q := &Queue{root}

	for !q.isEmpty() {
		// For every level reset sum = 0
		sum = 0


		// Iterate through the queue for each level to get the sum at that level
		for i := 0; i < len(*q); i++ {
			pop := q.Pop()
			sum += pop

			if pop.Left != nil {
				q.Push(pop.Left)
			}

			if pop.Right != nil {
				q.Push(pop.Right)
			}
		}
	}

	return sum
}