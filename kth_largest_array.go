/**
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
**/


import "container/heap"

type Pqueue []int

func (q *Pqueue) Push(x interface{}) { 
    *q = append(*q, x.(int))
}

func (q *Pqueue) Pop() interface{} {
  l := len(*q)
  popped := (*q)[l-1]
  *q = (*q)[0:l-1]
  return popped
}

func (q Pqueue) Len() int { 
  return len(q) 
}

func (q Pqueue) Less(i, j int) bool{ 
  return q[i] < q[j] 
}
func (q Pqueue) Swap(i, j int) { 
  q[i],q[j] = q[j],q[i] 
}


func findKthLargest(nums []int, k int) int {
  l := len(nums)
  if l ==0 {
    return 0
  }
  
  if l == 1 {
    return nums[0]
  }
  
  pq := &Pqueue{}
  
  for i := 0; i < l; i++ {
    heap.Push(pq, nums[i])
      
      if len(*pq) > k {
          heap.Pop(pq)
      }
  }

    return (*pq)[0]
    
  
}