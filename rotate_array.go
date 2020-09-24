/**
Given an array, rotate the array to the right by k steps, where k is non-negative.

Follow up:

Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
Could you do it in-place with O(1) extra space?
 

Example 1:

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]



[1,2,3,4,5,6,7] i = 6
[1,2,3,4,5,5,6] = 5
[1,2,3,4,4,5,6] = 4
[1,2,3,3,4,5,6] = 3
[1,2,2,3,4,5,6] = 2

[1,1,2,3,4,5,6] = 1

**/

package main

import "fmt"

func rotateArray(arr []int, k int) {
	l := len(arr)
	if l == 0  || l==1{
		return
	}

	for k > 0 {
		temp := arr[l-1]
		for i := l-1; i >0; i-- {
			arr[i] = arr[i-1]
		}
		arr[0] = temp
		k--
	}
}


func main() {
	rotateArray([]int{1}, 3)
}