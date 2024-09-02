package main

import "fmt"

func majorityElement(nums []int) int{
	candidate, count := 0, 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
			count = 1
		} else if num == candidate {
			candidate = num
			count++
		} else {
			count--
		}
	}

	return candidate
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}