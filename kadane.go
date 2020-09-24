package main 

import (
	"fmt"
)

func maxSubArray(A []int )  (int) {
    if len(A) == 0 {
        return 0
    }
    
    if len(A) == 1  {
        return A[0]
    } 
    
    global_sum := A[0]
    max_sum_in_segment := A[0]
    
    for i := 1; i <= len(A) - 1; i++ {
        num := A[i]
        fmt.Println(num)
        if num > num + max_sum_in_segment {
            max_sum_in_segment =  num 
        } else {
            max_sum_in_segment = num + max_sum_in_segment
        }
        
        if max_sum_in_segment > global_sum {
            global_sum = max_sum_in_segment
        }
    }
    
    return global_sum
}


func main () {
	arr := []int{-163, -20}
	fmt.Println(maxSubArray(arr))
}