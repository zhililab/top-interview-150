package main

import "fmt"

// canJump 判断是否可以到达数组的最后一个位置 https://leetcode.cn/problems/jump-game
func canJump(nums []int) bool {
    maxReach := 0 // 初始化当前能够到达的最远位置为 0
    
    // 遍历数组中的每一个元素
    for i, num := range nums {
        if i > maxReach {
            // 如果当前索引 i 超过了 maxReach，说明无法到达该位置，返回 false
            return false
        }
        // 更新 maxReach 为当前位置 i 加上该位置能够跳跃的最大步数 num
        if i + num > maxReach {
            maxReach = i + num
        }
        // 如果 maxReach 已经大于等于数组的最后一个索引，返回 true
        if maxReach >= len(nums) - 1 {
            return true
        }
    }
    // 遍历结束后，如果 maxReach 仍然小于最后一个索引，返回 false
    return false
}

func main() {
    nums1 := []int{2, 3, 1, 1, 4}
    fmt.Println(canJump(nums1)) // 输出 true

    nums2 := []int{3, 2, 1, 0, 4}
    fmt.Println(canJump(nums2)) // 输出 false
}