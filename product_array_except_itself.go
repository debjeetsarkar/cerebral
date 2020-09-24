/**
Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].

Example:

Input:  [1,2,3,4]

[6,2,3,4]

[1, 6, 12, 36]
[24 12 4 1]

[24, 72, 48, 36]

Output: [24,12,8,6]
Constraint: It's guaranteed that the product of the elements of any prefix or suffix of the array (including the whole array) fits in a 32 bit integer.

Note: Please solve it without division and in O(n).

Impl

1. Create an array to store the product of elements on the left of an element.
2. Create an array to store the product of elements on the right of the element.
3. Iterate through both the arrays and find product of both and return result.


left array : [1, 1, 2, 6]
right array : [24, 12, 4, 1]


initialise left_product = 1

Left array generation
	* Iterate from l -> r starting from 1, for every element, its arr[i-1] * left_product
	* push to left array

initialise right_product = 1

right array generation
	* Iterate from r -> l starting from l-1, for every element, its arr[j] * left_product
	* push to left array

**/

package main

import "fmt"

func prepend(array []int, value int) []int {
	l := len(array)
	if l ==0 {
		return append(array, value)
	}

	return append([]int{value}, array...)

}

func productExceptSelf(nums []int) []int {
	l := len(nums)
	result := []int{}

	leftProductArr := []int{1}
	rightProductArr := []int{1}

	leftProduct, rightProduct := 1, 1

	for i := 1 ; i < l ; i++ {
		leftProduct *=  nums[i-1]
		leftProductArr = append(leftProductArr, leftProduct)
	}

	for j := l - 2; j >=0; j-- {
		rightProduct *=  nums[j+1]
		rightProductArr = prepend(rightProductArr, rightProduct)
	}

	for i := 0;  i< l; i++ {
		result = append(result, (leftProductArr[i] * rightProductArr[i]))
	}

	return result
}

func main () {
	fmt.Println(productExceptSelf([]int{1,2,3,4}))
}

