package main

import "fmt"

// https://leetcode.cn/problems/majority-element-ii/description/
func majorityElement(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }

    // 初始化两个候选人和他们的计数器
    candidate1, candidate2, count1, count2 := 0, 1, 0, 0

    // 候选人选择阶段
    for _, num := range nums {
        if num == candidate1 {  // 如果当前数是第一个候选人，增加计数
            count1++
        } else if num == candidate2 {  // 如果当前数是第二个候选人，增加计数
            count2++
        } else if count1 == 0 {  // 如果第一个候选人计数为0，替换为当前数
            candidate1, count1 = num, 1
        } else if count2 == 0 {  // 如果第二个候选人计数为0，替换为当前数
            candidate2, count2 = num, 1
        } else {  // 如果当前数不是两个候选人之一，两个候选人计数都减1
            count1--
            count2--
        }
    }

    // 验证阶段
    count1, count2 = 0, 0
    for _, num := range nums {
        if num == candidate1 {  // 统计第一个候选人的最终出现次数
            count1++
        } else if num == candidate2 {  // 统计第二个候选人的最终出现次数
            count2++
        }
    }

    // 检查候选人是否满足出现次数超过 n/3 的条件
    result := []int{}
    if count1 > len(nums)/3 {
        result = append(result, candidate1)
    }
    if count2 > len(nums)/3 {
        result = append(result, candidate2)
    }

    return result
}

func main() {
    nums := []int{3, 2, 3}
    fmt.Println(majorityElement(nums)) // 输出 [3]
}