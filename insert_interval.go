/**
Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]


curr = 1,5 

Example 2:

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10]


**/



package main

import "fmt"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func isLeftOf(a, b []int) bool {
	return a[1] < b[0]
}


func isMergeable(a, b []int) bool{
	return a[1] >= b[0]
}

func insert(intervals [][]int, newInterval []int) [][]int {
	l := len(intervals)

	lNewInterval := len(newInterval)

	if l == 0 && lNewInterval > 0 {
		return [][]int{newInterval}
	}

	result := [][]int{}

	currentInterval := newInterval

	for _, interval := range intervals {
		if isLeftOf(interval, currentInterval) {          
			result = append(result, interval)
		} else if isMergeable(currentInterval,interval) {
			currentInterval[1] = max(interval[1], currentInterval[1])
			currentInterval[0] = min(interval[0], currentInterval[0])
           
		} else {
			result = append(result, currentInterval)
			currentInterval = interval
		}
	}
    
    result = append(result, currentInterval)
    
	return result
}
