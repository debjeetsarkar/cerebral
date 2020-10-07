/**
Given a string s, return the longest palindromic substring in s.

 

Example 1:

Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
**/


func isPalindrome(s string, left, right int) bool {
	for left < right {
		if string(s[left]) != string(s[right]) {
			return false
		}
		left++
		right--
	}

	return true
}


func findPalindrome(s string) string {
	left := 0
	right := len(s)-1
	maxPalindromeLength := 0
	result := ""
	for left < right {
		if isPalindrome(s, left, right) {
			if (right - left) > maxPalindromeLength {
				result = s[left, right+1]
			}
		} else if isPalindrome(s, left+1, right) {
			if (right - left + 1) > maxPalindromeLength {
				result = s[left+1, right+1]
			}
		} else if isPalindrome(s, left, right-1) {
			if (right - 1 - left) > maxPalindromeLength {
				result = s[left, right]
			}
		}

		left++
		right--

	}

	return result
}