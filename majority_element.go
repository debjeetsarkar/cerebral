/**
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Input: [3,2,3]
Output: 3
**/
package main

import "fmt"

func majorityElement(arr []int) int {
	l := len(arr)

	freqMap := make(map[int]int)

	for i := 0; i < l; i++ {
		v, ok := freqMap[arr[i]]
		if ok {
			freqMap[arr[i]] = v + 1
		} else {
			freqMap[arr[i]] = 1
		}
	}

	max := 0
	maxKey := arr[0]

	for key, val := range freqMap {
		if val > max {
			max = val
			maxKey = key
		}
	}

	return maxKey
}


func main () {
	fmt.Println(majorityElement([]int{3,2,3}))
}