/**

There are a total of numCourses courses you have to take, labeled from 0 to numCourses-1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?

 

Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take. 
             To take course 1 you should have finished course 0. So it is possible.


**/

type Queue []int

func (q *Queue) push(val) {
	(*q) = append((*q), val)
}

func (q *Queue) poll() int {
	polled := (*q)[0]
	*q = (*q)[1:len(*q)]
	return polled
}

func (q *Queue) isNotEmpty() bool {
	return len(*q) > 0 
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	// Create adjacency list and also populate indegree array.
	adjList := make(map[int][]int)
	indegreeArray := [numCourses]int{}
	for _, prerequisite := range prerequisites {
		adjListEntry, ok := adjList[prerequisite[1]]
		if ok {
			adjListEntry = append(adjListEntry, prerequisite[0])
		} else {
			list := []int{prerequisite[0]}
			adjList[prerequisite[1]] = list
		}
		indegreeArray[prerequisite[0]]++
	}

	// Go through the indegree array and fill the Queue with all whose indegree value is zero and initialise count as 0.
	for node, indegreeVal := range indegreeArray {
		if indegreeVal == 0 {
			queue.push(node)
		}
	}

	count := 0

	/** Iterate through the queue while its not empty:
		- Pop an element. 
		- Every time an element is popped, check if the value is 0 in the indegree array, if yes increase count
		- Check if popped is present in the adjacency list
			- If no, continue
			- If yes, get the neighbours
				- For every neighbour, decrease the value in indegree. 
				- If the value of the neighbour in indegree decreses to 0, add to queue
	**/
	for queue.isNotEmpty() {
		poll := queue.poll()
		if indegreeArray[poll] == 0 {
			count++
		}

		neighbours, ok := adjList[poll]
		if !ok {
			continue
		} else {
			for _, neighbour := range neighbours {
				indegreeArray[neighbour]--
				if indegreeArray[neighbour] == 0 {
					queue.push(neighbour)
				}
			}
		}
	}

	return count == numCourses
}