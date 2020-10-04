/**
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
**/

package main

import (
	"fmt"
	"strconv"
)

func addBinary(a, b string) string {
	la := len(a)
	lb := len(b)

	if la == 0 {
		return b
	}

	if lb == 0 {
		return a
	}

	carry := 0
	sum := 0
	i, j := la - 1, lb -1
	result := ""

	for i >= 0 || j >=0 {
		sum = 0
		if i >= 0 {
			sum += int(byte(a[i]) - '0')
			i--
		}

		if j >= 0 {
			sum += int(byte(b[j]) - '0')
			j--
		}

		sum += carry
		
		result = string((sum % 2) + '0') + result
		carry = sum / 2
	}

	if carry > 0 {
		result = string(carry + '0') + result
	}

	return result
}

func main () {
	fmt.Println(addBinary("11", "1"))
}