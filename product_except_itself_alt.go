/**


left array : [1, 1, 2, 6]
right array : [24, 12, 4, 1]

**/


func productExceptSelf(nums []int) []int {
	l := len(nums)

	if l == 0 {
		return []int{}
	}

	if l == 1 {
		return nums
	}

    result := make([]int, l)
    result[0] = 1
	temp := 1

	for i := 1; i < l; i ++ {
		result[i] = result[i-1] * nums[i-1]
	}

	for i := l-1 ; i >=0; i-- {
		result[i] *= temp
		temp *= nums[i]
	}

	return result
}

