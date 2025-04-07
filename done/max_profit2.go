package main

import "fmt"

func maxProfit(prices []int) int {
	i := 0
	profit := 0
	arrLen := len(prices)
	for i < arrLen-1 {
		currPrice := prices[i]
		if prices[i+1] > currPrice {
			profit += prices[i+1] - currPrice
		}
		i++
	}
	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
