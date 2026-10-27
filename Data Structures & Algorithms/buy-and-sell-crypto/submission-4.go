func maxProfit(prices []int) int {
	i,j := 0, 1
	maxP := 0
	for j < len(prices) {
		if prices[i] < prices[j] {
			maxP = max(maxP, prices[j] - prices[i])
		} else {
			i = j
		}
		j++
	}
	return maxP
}