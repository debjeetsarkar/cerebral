/**
Given a non-empty array of integers, return the k most frequent elements.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]
Note:

You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
It's guaranteed that the answer is unique, in other words the set of the top k frequent elements is unique.
You can return the answer in any order.
**/


import "container/heap"

type Item struct {
  value int
  freq int
}

type Pqueue []*Item

func (q Pqueue) Len() int { return len(q) }

func (q *Pqueue) Push(val interface{}) {
  item := val.(*Item)
  *q = append(*q, item)
}

func (q *Pqueue) Pop() interface{} {
  old := *q
  n := len(old)
  item := old[n-1]
  old[n-1] = nil  
  *q = old[0 : n-1]
  return item
}

func (q Pqueue) Swap(i, j int) {
  q[i], q[j] = q[j], q[i]
}

func (q Pqueue) Less(i, j int) bool { return q[i].freq < q[j].freq }

func topKFrequent(nums []int, k int) []int {
  l := len(nums)
  result := []int{}

  if l == 0 {
    return result
  }

  if l == 1 {
    return nums
  }
  
  freqMap := make(map[int]int)

  for i := 0; i < l; i ++ {
    v, ok := freqMap[nums[i]]
    
    if ok {
      freqMap[nums[i]] = v + 1  
    } else {
      freqMap[nums[i]] = 1
    }
  }
  
  pq := &Pqueue{}

  for key, value := range freqMap {  
    heap.Push(pq, &Item{
      value : key,
      freq: value,
    })
      
      if len(*pq) > k {
          heap.Pop(pq)
      }
  }
  
    for k > 0 {
        result = append(result, heap.Pop(pq).(*Item).value)
        k--
    }
    
  
    return result
}