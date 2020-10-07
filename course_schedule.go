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

func (q *Queue) push(val int) {
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

    // Create the adjacency map

	adjMap := make(map[int][]int)
    indegreeArray := make([]int, numCourses)

	for _,prerequisite := range prerequisites {
		adjMapVal, ok := adjMap[prerequisite[1]]
		if ok {
			adjMap[prerequisite[1]] = append(adjMapVal, prerequisite[0])
		} else {
            list := []int{prerequisite[0]}
			adjMap[prerequisite[1]] = list
		}
		// Fill the indegree array
		indegreeArray[prerequisite[0]]++
	}


	// Topological sort queue
	q := &Queue{}

	// Loop over the indegree Array and for every zero value, push to the queue
	for course, indegreeVal := range indegreeArray {
		if indegreeVal == 0 {
			q.push(course)
		}
	}

	count := 0

	for q.isNotEmpty() {
		pop := q.poll()

		if indegreeArray[pop] == 0 {
			count++
		}

		// check for adjacent ones and reduce indegree value

		adjMapVal, ok := adjMap[pop]
		if ok {
			for _, neighbour := range adjMapVal {
				indegreeArray[neighbour]--
				if indegreeArray[neighbour] == 0 {
					q.push(neighbour)
				}
			}
        }
	}

	return count == numCourses
}