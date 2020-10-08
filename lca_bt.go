/**

Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

Given the following binary tree:  root = [3,5,1,6,2,0,8,null,null,7,4]


Example 1:

Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
Output: 3
Explanation: The LCA of nodes 5 and 1 is 3.

Approach

Use BFS

1. Push the root to the Queue. Set the parent to nil
2. If Queue is not empty loop over and found map's length < 2
 2.1 Pop from Queue
 2.2 If not in visited map 
  2.2.1 Check if popped is either p or q 
   2.2.1.1 If yes add, to found map.

 2.3 Check adjacent nodes

 	2.3.1 if node.Left is not in visited
 	 2.3.1.1 Add to the queue
 	 2.3.1.2 Add parent info to the visited map

    2.3.2 if node.Right is not in visited
 	 2.3.2.1 Add to the queue
 	 2.3.2.2 Add parent info to the visited map
**/

package main

import "fmt"

type Queue []*TreeNode


func (q *Queue) poll() *TreeNode {
	popped := (*q)[0]
	(*q) = (*q)[1:len(*q)]

	return popped
}

func (q *Queue) push(node *TreeNode) {
	(*q) = append((*q), node)
}

func (q *Queue) isEmpty() bool {
	return len(*q) < 1
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	bfsQueue := &Queue{}
	bfsQueue.push(root)
	parentMap := make(map[*TreeNode]*TreeNode)
	found := 0

	for !bfsQueue.isEmpty() && found < 2 {
		popped := bfsQueue.poll()

		if popped.Val == p.Val || popped.Val == q.Val {
			found++
		}

		if popped.Left != nil {
			parentMap[popped.Left] = popped
            bfsQueue.push(popped.Left)
		}

		if popped.Right != nil {
			parentMap[popped.Right] = popped
            bfsQueue.push(popped.Right)
		}
	}

	ancestorMap := make(map[*TreeNode]*TreeNode)
    
    
	for p != nil {
		ancestorMap[p] = parentMap[p]
		p = parentMap[p]
	}
    
    
	for q != nil {
		_, ok := ancestorMap[q]
		if ok {
			return q
		}

		q = parentMap[q]
	}

	return nil
}




