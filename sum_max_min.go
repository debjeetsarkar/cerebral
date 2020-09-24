/**
Problem Description

Given an array A of size N. You need to find the sum of Maximum and Minimum element in the given array.

NOTE: You should make minimum number of comparisons.

Problem Constraints
1 <= N <= 105

-109 <= A[i] <= 109


Input Format
First and only argument is an integer array A of size N.

Output Format
Return an integer denoting the sum Maximum and Minimum element in the given array.



Example Input
Input 1:

 A = [-2, 1, -4, 5, 3]
Input 2:

 A = [1, 3, 4, 1]


 -109 ----------- -4 -2 -- 0 -1--3--5 ------- 109

if its a positive number then 109 - num will be minimum
if its a negative number then -109 - (-4) = -105 and -109 - (-2) = -107 will be minium

 A = [-2, 1, -4, 5, 3]

Iterating through the array if num is negative then
 min = min(min, -109 - (num))
 and max max(max, -109 - (num)
else
 max = min(max, 109- num)

 At the end, sum the min + max and return


 -2: max(-109, -109 - (-2)) : min => -107 current min = -2
 -4: max(-107, -109 - (-4)) : min => -105 current min = -4 

 1: min(109, 109-1) = 108, current max = 1
 5: min (108, 109- 5) = 104, current max = 5
 3: min (104, 109-3) = 106, still 104


return -4 + 5 = 1

**/


package main 

import (
	"fmt"
)

func solve(A []int )  (int) {
	l := len(A)

	if l == 0 {
		return 0
	}

	if l == 1 {
		return A[0]
	}

	if l == 2  {
		return A[1] + A[0]
	}

	max := A[0]
	min := A[0]

	for i :=1; i < l; i++ {
		if A[i] <= min {
			min = A[i]
		}

		if A[i] >= max {
			max = A[i]
		}
	}
	return max + min
}

func main() {
	fmt.Println(solve([]int{-2, 1, -4, 5, 3}))
}



