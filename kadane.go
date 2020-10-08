package main 

import "math"

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxSubArray(nums []int) int {    
    maxSegmentSum := 0
    globalSum := math.MinInt32
    
    for _, num := range nums {
        maxSegmentSum = max(num, num + maxSegmentSum)
        globalSum = max(globalSum, maxSegmentSum)
    }
    
    return globalSum
}

func main () {
	arr := []int{-163, -20}
	fmt.Println(maxSubArray(arr))
}