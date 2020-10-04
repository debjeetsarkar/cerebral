/**
Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.

Note:


The length of both num1 and num2 is < 5100.
Both num1 and num2 contains only digits 0-9.
Both num1 and num2 does not contain any leading zero.
You must not use any built-in BigInteger library or convert the inputs to integer directly.
**/


package main 

import (
	"fmt"
)


func addStrings(num1 string, num2 string) string {
	l1 := len(num1) -1 
	l2 := len(num2) -1

	if l1 == -1 {
		return num2
	}

	if l2 == -1 {
		return num1
	}

	sum := 0
	str := ""
	carry := 0

	for l1 >= 0 || l2 >= 0 {
		val1 := 0
		val2 := 0
		if l1>=0 && byte(num1[l1]) - '0' >= 0 {
			val1 = int(byte(num1[l1]) - '0')
		}

		if l2 >=0 && byte(num2[l2]) - '0' >= 0 {
			val2 = int(byte(num2[l2]) - '0')
		}

		sum = val1 + val2 + carry

		if sum > 9 {
			carry = sum / 10
			str = string((sum % 10) + '0') + str

		} else {
			carry = 0
			str = string(sum + '0') + str
		} 

		l1 --
		l2 --
	}

	if carry > 0 {
		str = string(carry + '0') + str
	}

	return str
}

func main () {
	fmt.Println(addStrings("9", "99"))
}