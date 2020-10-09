/**
Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:

Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"

If there is no such window in S that covers all characters in T, return the empty string "".
If there is such window, you are guaranteed that there will always be only one unique minimum window in S.

Impl:

1. Create the window map of T with count of every characters.
2. Set counter to number of unique characters.

3. Iterate through S and for every character check if in the window map. 
 3.1 For every ocuurance decrease value in map by 1.
 3.2 For a given character in map, when value becomes zero after decrease, decrease counter by 1.

4. When counter is 0 then we have a substring that has atleast all the characters in T.
 4.1 Till counter is 0 we try to decrease the size of the window by checking if removing one character from window.
 4.2 For doing this we maintain a start variable init at 0 and we check if value at start is present in window map.
 	4.2.1 If yes, we increase its value in the map.
 	4.2.2 If value after increase is > 0 then we increase the counter, this breaks this loop.
 		4.2.2.1 Calculate if slider - start < min_window, then push back data in result
 4.3 Increase value of start for sliding the window.

**/

package main 

import "fmt"

func minWindow(s, t string) string {
	l := len (s)
	result := ""

	if l == 0 {
		return result
	}

	start := 0 

	window := make(map[string]int)
	counter := 0
	min_window := l

	for i := 0; i < len(t); i++ {
		val, ok := window[string(t[i])]
		if ok {
			window[string(t[i])] = val + 1
		} else {
			window[string(t[i])] = 1
			counter++
		}
	}

	for slider := 0; slider <l ; slider++ {
		v, ok := window[string(s[slider])]
		if ok {
			if v == 1 {
				counter--
			}
			window[string(s[slider])] = v - 1
		}

		for counter == 0 {
			startChar := string(s[start])

			value, ok := window[startChar]
			if ok {
				window[startChar] = value + 1
				if window[startChar] > 0  {
					if slider - start + 1 <= min_window {
						min_window = slider - start + 1
						result = s[start:slider + 1]
					}
					counter++
				}
			}

			start++
		}
	}

	return result
}


func main() {
	fmt.Println(minWindow("a", "a"))
}