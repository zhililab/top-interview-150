package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/candy
func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}

	// 每个孩子至少分一个糖果
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}

	// 从左到右遍历
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// 从右到左遍历
	for i := n-2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = int(math.Max(float64(candies[i]), float64(candies[i+1]+1)))
		}
	}

	totalCandies := 0
	for _, candy := range candies {
		totalCandies += candy
	}

	return totalCandies
}

func main() {
	ratings := []int{1, 0, 2}
    fmt.Println(candy(ratings)) // 输出 5

    ratings = []int{1, 2, 2}
    fmt.Println(candy(ratings)) // 输出 4
}