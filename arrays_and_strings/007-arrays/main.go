package main

import "fmt"

// 反转数组的指定部分
func reverse(nums []int, start int, end int) {
    for start < end {
        nums[start], nums[end] = nums[end], nums[start]
        start++
        end--
    }
}

// 旋转数组
func rotate(nums []int, k int) {
    n := len(nums)
    k = k % n  // 如果 k 大于数组长度，则取模
    reverse(nums, 0, n-1)    // 反转整个数组
    reverse(nums, 0, k-1)    // 反转前 k 个元素
    reverse(nums, k, n-1)    // 反转剩余的元素
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6, 7}
    k := 3
    rotate(nums, k)
    fmt.Println(nums)  // 输出 [5, 6, 7, 1, 2, 3, 4]
}