package main

import "fmt"

type Node struct {
	item int
}

type Stack struct {
	items [] *Node
	count int
}

func NewStack() *Stack{
	return &Stack{}
}

func (n *Node) String() string {
	return fmt.Sprint(n.item)
}


func(s *Stack) Push(val *Node) {
 s.items = append(s.items, val)
 s.count++
}

func(s *Stack) Pop() *Node{
	l := len(s.items)
	if l == 0 { return nil }

	n := s.items[l-1]
	s.items = s.items[: l -1]
	s.count--
	return n
}

func main() {
	s := NewStack()
	s.Push(&Node{3})
	s.Push(&Node{5})
	s.Push(&Node{7})
	s.Push(&Node{9})
	fmt.Println(s.Pop(), s.Pop(), s.Pop(), s.Pop())
}