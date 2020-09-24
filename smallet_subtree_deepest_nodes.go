/**
Given the root of a binary tree, the depth of each node is the shortest distance to the root.

Return the smallest subtree such that it contains all the deepest nodes in the original tree.

A node is called the deepest if it has the largest depth possible among any node in the entire tree.

The subtree of a node is tree consisting of that node, plus the set of all descendants of that node.


Input: root = [3,5,1,6,2,0,8,null,null,7,4]
Output: [2,7,4]
Explanation: We return the node with value 2, colored in yellow in the diagram.
The nodes coloured in blue are the deepest nodes of the tree.
Notice that nodes 5, 3 and 2 contain the deepest nodes in the tree but node 2 is the smallest subtree among them, so we return it.


**/



package main

import "fmt"


type TreeNode struct {
  Val int
  Left *TreeNode
  Right *TreeNode
}

type Queue []*TreeNode

func(q *Queue) Push(val *TreeNode) {
  *q = append((*q), val)
}

func(q *Queue) Pop() *TreeNode {
  popped := (*q)[0]
  *q = (*q)[1:len(*q)]
  return popped
}

func(q *Queue) isEmpty() bool{
  return len(*q) < 1
}  

func smallestSubtree(root *TreeNode) *TreeNode {
 /** 
  1. Perform BFS on the tree and get the deepest level nodes.
    1.1. During this process, populate a parent map <node : node's parent>
    1.2. During this process, for each level, capture the nodes in a list
    1.3. End of this process, we have all the parents populated and the deepest nodes in the list
    

  2. After the BFS process in 1.
    2.1 Create a set and loop over the items in the list and put their parents in the set.
    2.2 Continue the loop till the size of the set is > 1, when 1, we have got out our root of the subtree.
    2.3 Return the value from the set
    
    TC : O(n)
    SC : O(n)
  */
  
  lca := nil
  parentMap := make(map[*TreeNode]*TreeNode)
  queue := &Queue{}
  
  if root == nil {
    return lca 
  }
  
  parentMap[root] = nil
  queue.Push(root)
  leaves := []*TreeNode{}
  
  for !queue.isEmpty() {
    leaves = nil
    
    l := len(queue)
    for i := 0; i < l; i++ {
      popped := queue.Pop()
      leaves = append(leaves, popped)
      
      if popped.Right != nil {
        parentMap[popped.Right] = popped
        queue.Push(popped.Right)
      }
      if popped.Left != nil {
        parentMap[popped.Left] = popped
        queue.Push(popped.Left)
      }
    }
  }
  
  
  if len(leaves) == 1 {
   return leaves[0]
  }
  
  ancestorSet := make(map[*TreeNode]int)
  
  for _, leaf := range leaves {
    ancestorSet[leaf] = 0
  }
  
  for len(ancestorSet) >= 1 {
    for key, _ := range ancestorSet {
      if len(ancestorSet) == 1 {
        return key
      }
      delete(ancestorSet, key)
      ancestorSet[parentMap[key]] = 0
    }
  }
  
  return nil
}
