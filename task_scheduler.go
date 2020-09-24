/**
Given a characters array tasks, representing the tasks a CPU needs to do, where each letter represents a different task. Tasks could be done in any order. Each task is done in one unit of time. For each unit of time, the CPU could complete either one task or just be idle.

However, there is a non-negative integer n that represents the cooldown period between two same tasks (the same letter in the array), that is that there must be at least n units of time between any two same tasks.

Return the least number of units of times that the CPU will take to finish all the given tasks.


Input: tasks = ["A","A","A","B","B","B"], n = 2
Output: 8
Explanation: 
A -> B -> idle -> A -> B -> idle -> A -> B
There is at least 2 units of time between any two same tasks.

**/

package main

import "container/heap"

type PriorityQueue []int

func (p PriorityQueue) len() {
	return len(p)
}

func (p PriorityQueue) isEmpty() {
	return len(p) < 0
}

func (p PriorityQueue) Less(i, j int) {
	return len(p[j]) > len(p[i])
}

func (p PriorityQueue) swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(val interface{}) {
	*h = append(*h, val.(int))
}

func (p *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


func leastInterval(tasks []byte, n int) int {
    // Create priority queue.
    pq := PriorityQueue{}

    // Iterate through tasks and put the frequency in a map
    freqMap := make(map[byte]int)
    for _, task := range tasks {
    	value, ok := freqMap[byte(task)];
    	if ok {
    		freqMap[byte(task)] = value + 1 
    	} else {
    		freqMap[byte(task)] = 1
    	}
    }

    for _,val := range freqMap {
    	heap.Push(&pq, val)
    }

    /**
    Till the time the heap is not empty:
    	- Take loop of 0 -> n+1 and if the heap is not empty then
    	 	- pop from the heap and add to a list, processing list 
    	- Loop through processing list and for every value
    		- Check if the decreased value is > 0
    		- If yes, put back to the heap
    	Keep adding to result, 
    		- If heap is empty then processing list's length or else n+1 (cooldown period)

    **/
    result := 0
    for !pq.isEmpty() {
    	processingList := []int
    	for i :=0; i < n+1; i++ {
    		if !pq.isEmpty() {
    			processingList = append(processingList, (heap.Pop(&pq)).(int))
    		}
    	}

    	for _,task := range processingList {
    		if task > 1 {
    			heap.Push(&pq, task - 1)
    		}
    	}

    	if pq.isEmpty() {
    		result += len(processingList)
    	} else {
    		result += n+1
    	}
    }

    return result
}