package main

import "fmt"

type Stk [] string

func(s *Stk) Push(val string) {
 *s = append(*s, val)
}

func(s *Stk) Pop() string{
	l := len(*s)
	if l == 0 { return "" }

	n := (*s)[l-1]
	*s = (*s)[: l -1]
	return n
}

func main() {
	s := &Stk{}
	s.Push("(")
	s.Push("[")
	s.Push("]")
	s.Push(")")
	fmt.Println(s.Pop(), s.Pop(), s.Pop(), s.Pop())
}