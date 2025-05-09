package main

func maxProfitII(prices []int) int {
	profit := 0
	for i := 1; i < len(prices)-1; i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
