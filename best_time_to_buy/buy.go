package buy

func MaxProfit(prices []int) int {
	// we are given an array of prices
	// and we need to buy and sell
	// we can only hold one stock at a time
	// so this problem boils down to
	// partitioning the array into ascending sub arrays and buying
	// and selling at each end point
	profit := 0
	startHead := 0
	endHead := 0
	for i := 0; i < len(prices)-1; i++ {
		currentValue := prices[i]
		nextValue := prices[i+1]
		if nextValue < currentValue {
			// price goes down, sell what we have
			valueIncrease := prices[endHead] - prices[startHead]
			profit += valueIncrease

			startHead = i + 1
			endHead = i + 1
		} else {
			// price goes up or stays the same
			// keep extending the subarray
			endHead = i + 1
		}
	}
	valueIncrease := prices[endHead] - prices[startHead]
	profit += valueIncrease

	return profit
}
