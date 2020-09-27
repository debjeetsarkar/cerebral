/**
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
**/

package main

import "fmt"


func findStartingIndex(nums []int, target int) int {
	low, high, index := 0, len(nums) -1 , -1

	for low <= high {
		mid := low + (high-low) / 2

		if nums[mid] == target {
			index = mid
		}

		if nums[mid] >= target {
			high = mid -1
		} else {
			low = mid + 1
		}
	}

	return index
}

func findEndingIndex(nums []int, target int) int {
	low, high, index := 0, len(nums) -1, -1

	for low <= high {
		mid := low + (high-low) / 2

		if nums[mid] == target {
			index = mid
		}

		if nums[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return index
}

func searchRange(nums []int, target int) []int {
	l := len(nums)

	if l == 0 {
		return []int{-1, -1}
	}

	if l == 1 {
		if nums[0] == target {
			return []int{0,0}	
		}
		return []int{-1, -1}	
	}

	res := []int{}

	res = append(res, findStartingIndex(nums, target), findEndingIndex(nums, target))

	return res
}


func main () {
	fmt.Println(firstLast([]int{5,7,7,8,8,10}, 8))
}