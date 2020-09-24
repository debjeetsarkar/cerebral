/**
283. Move Zeroes

Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.


[10,1,0,3,12]
[10,1,3,0,12]

[0,1,0,3,12]


start = 0
i = 1

arr[start] * arr[i] == 0 && arr[i] != 0

arr[start] = arr[i]
arr[i] = 0
start++

**/

package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	l := len(nums)

	if l == 0 || l == 1{
		return
	}
	start := 0

	for i := 1; i <l; i ++ {

		if nums[i] == nums[start] && nums[start] == 0 {
			continue
		}

		if nums[start] - nums[i] == -1 * nums[i] {
			nums[start] = nums[i]
			nums[i] = 0
			start++
		} else {
			start++
		}
	}
}

func main() {
	arr := []int{0,1,0}
	moveZeroes(arr)

	fmt.Println(arr)
}