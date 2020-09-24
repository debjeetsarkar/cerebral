/**
66. Plus One

Given a non-empty array of digits representing a non-negative integer, increment one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contains a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

 
Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.


[9]
**/

package main

import "fmt"


func prepend(arr []int, val int) []int{
	return append([]int{val}, arr...)
}

func increament(digits []int) []int {
	l := len(digits)

	if l == 0 {
		return digits
	}

	carry := 0
	index := l-1
	for index >= 0 {
		var incrementedVal int
		if index == l-1  {
			incrementedVal = digits[index] +1
		} else {
			incrementedVal = carry + digits[index]	
		}
		
		if incrementedVal <= 9 {
			digits[index] = incrementedVal

			return digits
		} else {
			carry = incrementedVal / 10
			if index == 0 {
				digits[index] = 0
				digits = prepend(digits, carry)
			} else {
				digits[index] = 0
			}
		}

		index--
	}

	return digits
}


func main () {
	ip := []int{1,2,3}
	fmt.Println(increament(ip))
}