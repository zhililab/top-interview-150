package main

import "fmt"

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    // 慢指针 i，快指针 j
    i := 0
    for j := 1; j < len(nums); j++ {
        // 当 nums[j] 与 nums[i] 不同时，说明找到了新的元素
        if nums[j] != nums[i] {
            i++ // 慢指针 i 向前移动一位
            nums[i] = nums[j] // 将新元素放到慢指针位置
        }
    }
    return i + 1 // 返回数组新长度
}

func main() {
    nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
    newLength := removeDuplicates(nums)
    fmt.Println(nums[:newLength]) // 输出删除重复项后的数组
}