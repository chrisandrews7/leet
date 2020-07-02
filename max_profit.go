package main

func MaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	lowest := prices[0]
	highest := prices[1]
	maxPrice := 0

	for _, price := range prices {
		if highest < price {
			highest = price
		}

		if lowest > price {
			highest = price
			lowest = price
		}

		if max := highest - lowest; max > maxPrice {
			maxPrice = max
		}
	}

	return maxPrice
}
