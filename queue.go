package main

import "fmt"

type Queue []string

func (q *Queue) push(str string) {
	*q = append(*q, str)
}

func (q *Queue) pop() string{
	if len(*q) == 0 {
		return ""
	}
	popped := (*q)[0]
	*q = (*q)[1: len(*q)]
	return popped
}

func main () {
	myQueue := &Queue{}
	fmt.Println("Popped element: ", myQueue.pop())
	myQueue.push("1")
	myQueue.push("2")
	myQueue.push("3")
	fmt.Println("Queue: ", myQueue)
	fmt.Println("Popped element: ", myQueue.pop())
	fmt.Println("Popped element: ", myQueue.pop())
	fmt.Println("Popped element: ", myQueue.pop())
}