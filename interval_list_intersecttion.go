/**

Given two lists of closed intervals, each list of intervals is pairwise disjoint and in sorted order.

Return the intersection of these two interval lists.

(Formally, a closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.  The intersection of two closed intervals is a set of real numbers that is either empty, or can be represented as a closed interval.  For example, the intersection of [1, 3] and [2, 4] is [2, 3].)

Example: 

Input: 
A = [[0,2],[5,10],[13,23],[24,25]] 
B = [[1,5],[8,12],[15,24],[25,26]]
Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]

 
**/
package main

import "fmt"

ffunc max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}


func intervalIntersection(A, B [][]int) [][]int {
	al := len(A)
	bl := len(B)

	result := [][]int{}

	if al == 0 || bl == 0 {
		return result
	}

	pointerA := 0
	pointerB := 0

	for pointerA < al && pointerB < bl {
        
        // if A[pointerA][0] > B[pointerB][1] {
        //     pointerB++
        //     continue
        // } else if A[pointerA][1] < B[pointerB][0] {
        //     pointerA++
        //     continue
        // }
		startingInterval, endingInterval := max(A[pointerA][0], B[pointerB][0]), min(A[pointerA][1], B[pointerB][1])

		if startingInterval <= endingInterval {
			result = append(result, []int{startingInterval, endingInterval})
		}

		if A[pointerA][1] < B[pointerB][1] {
			pointerA++
		} else {
			pointerB++
		}
	}

	return result
}


func main() {
	fmt.Println(intervalIntersection())
}