/**
Best Time to Buy and Sell Stock

Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.

Approach

1. Keep 2 pointers, start and slider
2. Iterate over prices and for every value
	2.1 Check if slider < start , then start = slider
	2.2 If slider >= start, then assign max_profit = max(max_profit, slider-start)
3. Return max_profit

**/

package main

import "fmt"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	maxProfit := 0
	l := len(prices)

	if l == 0 {
		return maxProfit
	}

	minPriceDay := 0	
	for day, price := range prices {
		minPrice := prices[minPriceDay]
		if price < minPrice {
			minPriceDay = day
		} else if price >= minPrice {
			maxProfit = max(maxProfit, price - prices[minPriceDay])
		}
	}

	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
}