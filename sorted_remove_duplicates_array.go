/**
Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example 1:

Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
Example 2:

Given nums = [0,0,1,1,1,2,2,3,3,4],

[0,0,1,1,1,2,2,3,3,4]

new_len and one slider 
new_len = 0
slider = 1

arr[new_len] == arr[slider] and arr[slider + 1] > arr[new_len] :
	arr[new_len + 1] = arr[slider + 1]
	new_len = new_len + 1

arr = [0,1,1,1,1,2,2,3,3,4]
new_len = 1
slider = 2 


Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.

[0,0,1,1,1,2,2,3,3,4] 

[1, 2, 2]
**/

package main 

import (
	"fmt"
)

func removeDuplicates(arr []int) int {
  l := len(arr)
  if l ==0 {
    return l
  }
  
  start := 0
  for i:=1; i <l; i++ {
    if arr[start] != arr[i] {
      arr[start+1] = arr[i]
      start++
    }
  }
  
  return start +1
}

func main () {
	arr := []int{0,0,1,1,1,2,2,3,3,4}
	removeDup(arr)
	// fmt.Println(removeDup(arr))
	// fmt.Println(arr)
}