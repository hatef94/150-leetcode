package main

import "math"

func maxProfit(prices []int) int {
	profit, minn := 0, prices[0]
	for _, price := range prices[1:] {
		minn = int(math.Min(float64(minn), float64(price)))
		profit = int(math.Max(float64(profit), float64(price-minn)))
	}
	if profit < 0 {
		return 0
	}
	return profit
}
