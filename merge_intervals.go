/**
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.

Approach

1. Sort the array so that all the starting slots are in order in the array. O(nLogn)
2. Create a result list and push the first interval and set it as the current interval.
3. Loop over intervals and for each 
 3.1 If yes, update the current interval's end to the end of the i+1th interval.
 3.2 else, push the ith interval to the result and set that as the current interval

**/


import (
	"sort"
)

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}


func merge(intervals [][]int) [][]int {
	l := len(intervals)
	if l == 0 || l ==1 {
		return intervals
	}

	result := [][]int{}

	sort.SliceStable(intervals, func(i,j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	currentInterval := intervals[0]
	append(result, currentInterval)

	for _, interval := range intervals {
		currentEnd := currentInterval[1]
		nextStart := interval[0]
		nextEnd := interval[1]

		if currentEnd >= nextStart {
			currentInterval[1] = max(nextEnd, currentEnd)
		} else {
			currentInterval = interval
			append(result, currentInterval) 
		}  
	}

	return result
}