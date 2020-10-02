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

func isLeftOfCurrentInterval(intervalA []int, intervalB []int) bool {
	return intervalA[1] < intervalB[0]
}

func isIntervalMerging(intervalA []int, intervalB []int) bool {
	return intervalA[1] >= intervalB[0]
}

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


func insert(intervals [][]int, newInterval []int) [][]int {

	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
    
	currentInterval := newInterval
    result := [][]int{}

	for _,interval := range intervals {
		if isLeftOfCurrentInterval(interval, currentInterval) {
			result = append(result, interval)
		} else if isIntervalMerging(currentInterval,interval) {
            currentInterval[0] = min(currentInterval[0], interval[0])
			currentInterval[1] = max(currentInterval[1], interval[1])
		} else {
			result = append(result, currentInterval)
			currentInterval = interval
		}
	}
    
    return append(result, currentInterval)
}
