/**
Given an array of integers nums, sort the array in increasing order based on the frequency of the values. If multiple values have the same frequency, sort them in decreasing order.

Return the sorted array.

Example 1:

Input: nums = [1,1,2,2,2,3]
Output: [3,1,1,2,2,2]
Explanation: '3' has a frequency of 1, '1' has a frequency of 2, and '2' has a frequency of 3.
Example 2:

Input: nums = [2,3,1,3,2]
Output: [1,3,3,2,2]
Explanation: '2' and '3' both have a frequency of 2, so they are sorted in decreasing order.
Example 3:

Input: nums = [-1,1,-6,4,5,-6,1,4,1]
Output: [5,-1,4,4,-6,-6,1,1,1]
**/


package main

import (
	"container/heap"
)

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

func (q Pqueue) Less(i, j int) bool {
	if  q[i].freq == q[j].freq {
		return q[i].value > q[j].value
	}
	return q[i].freq < q[j].freq 
}


func frequencySort(nums []int) []int {
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
  }


    for len(*pq) > 0 {
        popped := heap.Pop(pq)

        fmt.Println(popped.(*Item))
        
        for i:= 0 ; i < popped.(*Item).freq; i ++ {
            result = append(result, popped.(*Item).value)    
        }
    }

  
    return result
}

func main() {
	fmt.Println(frequencySort([]int{-1,1,-6,4,5,-6,1,4,1}))
}