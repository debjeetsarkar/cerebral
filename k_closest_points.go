/**
We have a list of points on the plane.  Find the K closest points to the origin (0, 0).

(Here, the distance between two points on a plane is the Euclidean distance.)

You may return the answer in any order.  The answer is guaranteed to be unique (except for the order that it is in.)

 

Example 1:

Input: points = [[1,3],[-2,2]], K = 1
Output: [[-2,2]]
Explanation: 
The distance between (1, 3) and the origin is sqrt(10).
The distance between (-2, 2) and the origin is sqrt(8).
Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
We only want the closest K = 1 points from the origin, so the answer is just [[-2,2]].
**/

package main 

import (
	"container/heap"
)


type Point struct {
	distance int
	value []int
	index int
}

type PQueue []*Point

func (p PQueue) Len() int {
	return len(p)
}

func (p PQueue) Less (a, b int) bool {
	return p[a].distance > p[b].distance
}

func (p PQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p *PQueue) Push(x interface{}) {
	n := len(*p)
	point := x.(*Point)
	point.index = n
	*p = append(*p, point)
}

func (p *PQueue) Pop() interface{} {
	old := *p
	n := len(old)
	point := old[n-1]
	old[n-1] = nil  // avoid memory leak
	point.index = -1 // for safety
	*p = old[0 : n-1]
	return point
}


func Kclosest(points [][]int, k int) [][]int {

	pq := PQueue{}

	for i, point := range points {
		 p := &Point{
		 	distance : (point[0] * point[0]) + (point[1] * point[1]),
		 	value: point,
		 	index : i,
		 }
		 
		 heap.Push(&pq, p)
		 if pq.Len() > k {
		 	heap.Pop(&pq)
		 }
	}

	res := make([][]int, k)

	for i := 0; i < len(pq); i++ {
		res[i] = pq[i].value
	}
	return res
}




func main () {
	fmt.Println(Kclosest([][]int{{1,3}, {-2,2}}, 1))
}