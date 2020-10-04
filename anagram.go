/**

Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false

0. If strings are of different length or both empty then return false.
1. Loop through the string1 and push it to a map with the count.
2. Loop through string2 and check if its present. If not return false.
 2.1. If present decrease count and
3. After loop if any of the keys have count !=0 then return false. else
	3.1 return true

**/


package main

import (
	"fmt"
)

func isAnagram(s1, s2 string) bool {
  l1 := len(s1)
  l2 := len(s2)

  if l1 != l2 {
  	return false
  } else {
  	if l1 == 0 {
  		return true
  	}
  }

  counterMap := make(map[string]int)
  for i := 0; i < l1; i ++ {
  	counterMap[string(s1[i])] = counterMap[string(s1[i])] + 1
  	counterMap[string(s2[i])] = counterMap[string(s2[i])] - 1 
  }

  for _, v := range counterMap {
  	if v !=0 {
  		return false
  	}
  }

  return true
}


func main () {
	fmt.Println("Result: ", isAnagram("anagram", "nagaram"));
}