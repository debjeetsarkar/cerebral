/**
Given an array of integers and an integer k, you need to find the total number of continuous subarrays whose sum equals to k.

Example 1:

Input:nums = [1,1,1,1], k = 3
Output: 2

[1,5,1,1], k = 2

**/
package main

import "fmt"

func subarraySum(nums []int , k int) int {
	result := 0
	l := len(nums)

	if l == 0 {
		return 0
	}

	con_sum := 0
	ref_map := make(map[int]int)

	for i := 0; i <l ; i ++ {
		con_sum += nums[i]

		if con_sum == k {
			result++
		}

		v, ok := ref_map[con_sum - k]
		if ok {
			result += v
		} 

		val, present := ref_map[con_sum]
		if present {
			ref_map[con_sum] = val + 1
		} else {
			ref_map[con_sum] = 1	
		}
	}

	return result
}

func main () {
	fmt.Println(subarraySum([]int{1,2,1,2,1}, 3))
}
