package main

import "fmt"

// jump 函数计算从第一个位置跳到最后一个位置的最少跳跃次数
func jump(nums []int) int {
    jumps, currentEnd, farthest := 0, 0, 0

    // 遍历每个位置
    for i := 0; i < len(nums)-1; i++ {
        // 更新当前能到达的最远位置
        farthest = max(farthest, i + nums[i])
        // 当到达当前跳跃范围的末尾时
        if i == currentEnd {
            jumps++
            currentEnd = farthest
        }
    }

    return jumps
}

// 辅助函数，返回两个整数中的较大值
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    nums := []int{2, 3, 1, 1, 4}
    fmt.Println(jump(nums)) // 输出 2
}