package main

func maxProfit(prices []int) int {
	maxProfit := 0
	minBuyPrice := prices[0]
	for _, sellPrice := range prices {
		maxProfit = max(maxProfit, sellPrice-minBuyPrice)
		minBuyPrice = min(minBuyPrice, sellPrice)
	}
	return maxProfit
}
