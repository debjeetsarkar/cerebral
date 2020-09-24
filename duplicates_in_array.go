package main

import (
	"fmt"
)

/**
Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements that appear twice in this array.

Could you do it without extra space and in O(n) runtime?

Example:
Input:
[4, 3, 2, 7, 8, 2, 3, 1]
[3, 2]


[4 4]
Output:
 */
func Abs(a int) int{
	if a < 0  { return -1 * a }
	return a
}

func findDuplicates(nums []int) []int {
 len := len(nums)
 var result []int
 if len == 0 { return result }
 if len == 1 { return result }

for index := 0; index < len; index++ {
	if Abs(nums[index]) < len {
		if nums[Abs(nums[index])] > 0 {
			nums[Abs(nums[index])] = -1 * nums[Abs(nums[index])]
		} else {
			result = append(result, nums[Abs(nums[index])])
		}
	}
}
 return result
}

func main () {
	inp := []int{4,3,2,7,8,2,3,1}
	fmt.Println(findDuplicates(inp))
}

