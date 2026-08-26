func maxProfit(prices []int) int {
	var maxProfit int = 0
	var minBuy int = math.MaxInt32
	for _, sell := range prices {
		if sell - minBuy > maxProfit {
			maxProfit = sell - minBuy
		}
		if sell < minBuy {
			minBuy = sell
		}
	}
	return maxProfit
}
