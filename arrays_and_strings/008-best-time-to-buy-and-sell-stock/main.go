package main 

import "fmt"

// 计算最大利润 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0
	
	for _, price := range prices {
		if price < minPrice {
			minPrice = price 
		} else if price - minPrice > maxProfit{
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))	// 输出 5
}