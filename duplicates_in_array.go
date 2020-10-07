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


[0, 0, 0, 0, 0, 2, 3, 1]


[4 4]
Output:
 */

func findDuplicates(nums []int) []int {
 len := len(nums)
 var result []int
 if len == 0 { return result }
 if len == 1 { return result }

for index, val := range nums {
	nums[val] = 0
}

for index, val := range nums {
	if val != 0 {
		result = append(result, val)
	}
}

return result

func main () {
	inp := []int{4,3,2,7,8,2,3,1}
	fmt.Println(findDuplicates(inp))
}

