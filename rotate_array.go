/**
Given an array, rotate the array to the right by k steps, where k is non-negative.

Follow up:

Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
Could you do it in-place with O(1) extra space?
 

Example 1:

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


func reverse(arr []int, start, end int) {
	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp

		start++
		end--
	}
}

func rotateArray(arr []int, k int) {
	l := len(arr)

	if l == 0 {
		return
	}

	k = k % len(arr)
	reverse(arr, 0, l-1)
	reverse(arr, 0, k-1)
	reverse(arr, k, l-1)
}


func main() {
	rotateArray([]int{1}, 3)
}