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


func reverse(s string) string {
	reversed := ""
	for _, val := range s {
		reversed =  string(val) + reversed
	}

	return reversed
}

func longestPalindrome(s string) string {
    l := len(s)

    result := ""

    if l == 0 {
    	return result
    }
}