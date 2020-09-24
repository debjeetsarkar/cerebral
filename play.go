package main

import (
	"fmt"
	"sort"
)



func demo(a string) int{
		l1 := len(a) -1

		if l1 > 0 {
			fmt.Println(byte(a[l1]) - '0')
		}

		return 1
}



func delChar(str string, index int) string {
	s := []rune(str)
    return string(append(s[0:index], s[index+1:]...))
}



func isLetter(r rune) bool {
  return ((r >= 'a' && r <= 'z') || (r >= '0' && r <= '9'))
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}



func main() {

	// type ListNode struct {
	// 	Val int
	// 	Next *ListNodex
	// }

	rangearray := [][]int{{1,3}, {5,8}, {2,6}}

	sort.SliceStable(rangearray, func(i, j int) bool {
        return rangearray[i][0] < rangearray[j][0]
	})


	for i := 0; i < len(rangearray); i++ {
		fmt.Println(rangearray[i][1])
	}

	// fmt.Println(max(9,10))


	// fmt.Println(isLetter('_'));

	// fmt.Println(delChar("a)b(c)d", 1))


	// result :=  &ListNode{ Val : 0, Next: nil}
	// ref := result

	// fmt.Println("Address result 1: ", result)
	// fmt.Println("Address ref 1: ", ref)
	// a := &ListNode{ Val: 10}
	// result.Next = a
	// result = result.Next
	// fmt.Println("Address result 2: ", result)
	// fmt.Println("Address ref 2: ", ref)
	// b := &ListNode{ Val: 10}
	// result.Next = b
	// result = result.Next
	// fmt.Println("Address result 3: ", result)
	// fmt.Println("Address ref 3: ", ref)

	// fmt.Println("Byte value: ", byte('z'))
}