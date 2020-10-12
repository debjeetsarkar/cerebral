/**
Given a non-empty array nums containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

 

Example 1:

Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].
Example 2:

Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.


Approach:

Edge cases:

Not possible if the array is empty
Not possible if the totalSum of elements in the array is odd


Array to be divided into two parts such that each part has equal subset sum that is totalSum/2
Use the totalSum/2 as the targetSum and apply that as the weigth of the knapsack

Create the knapsack matrix of size 
  rows: len of array + 1
  columns : weight of knapsack + 1

Initialise the matrix with column = 0 as true and rest as false

Apply knapsack algorithm by iterating over values over rows over columns
Return the last value of the knapsack matrix, that determines if the targetSum is possible or not to be acheived.

**/

func canPartion(nums []int) bool {
  l := len(nums)
  if l == 0 {
    return false
  }
  
  totalSum := 0
  for i:= 0; i< l; i++ {
    totalSum += nums[i]
  }
 
  if totalSum % 2 != 0 {
    return false
  }
  
  targetSum := totalSum / 2
  knapsack := make([][]bool, l + 1)
  
  for i:=0 ; i < len(knapsack); i++ {
    knapsack[i] = make([]bool, targetSum + 1)
  }
  
  for i:= 0; i < len(knapsack); i++ {
    for j = 0; j< len(knapsack[0]); j++ {
      if j == 0 {
        knapsack[i][j] = true
      } else {
        knapsack[i][j] = false
      }
    }
  }
  
  for i:= 1; i < len(knapsack); i++ {
    for j = 1; j< len(knapsack[0]); j++ {
      if (nums[i-1] <= j) {
        knapsack[i][j] = knapsack[i-1][j] || knapsack[i-1][j- nums[i-1]]
      } else {
        knapsack[i][j] = knapsack[i-1][j]
      }
    }
  }
  
  return knapsack[l][targetSum]
}