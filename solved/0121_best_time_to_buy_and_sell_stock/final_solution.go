package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	low := prices[0]
	maxP := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > low {
			if prices[i]-low > maxP {
				maxP = prices[i] - low
			}
		} else {
			low = prices[i]
		}
	}
	return maxP
}
