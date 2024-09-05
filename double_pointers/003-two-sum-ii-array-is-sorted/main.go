package main

import "fmt"

func twoSum(numbers []int, target int) []int {
    // 因为是非递减所有，所以可以使用双指针
    // 1. 定义 i, j := 0, len(numbers) - 1
    // 2. 
    // 2.1 开始遍历，如果 i + j == target ，则结束返回 [i+1, j+1]
    // 2.2 如果 i + j > target , 则j++ ，如果 i + j < target，则 i++

    i, j := 0, len(numbers) - 1
    for i < j {
        var sum = numbers[i] + numbers[j]
        if ( sum == target) {
            return []int{i + 1, j + 1}
        } else if (sum < target) {
            i++
        } else {
            j--
        }
    }

    return []int{-1, -1}
}

func main() {
    // 测试用例1
    numbers1 := []int{2, 7, 11, 15}
    target1 := 9
    fmt.Printf("示例1: 输入: numbers = %v, target = %d\n", numbers1, target1)
    fmt.Printf("输出: %v\n\n", twoSum(numbers1, target1))

    // 测试用例2
    numbers2 := []int{2, 3, 4}
    target2 := 6
    fmt.Printf("示例2: 输入: numbers = %v, target = %d\n", numbers2, target2)
    fmt.Printf("输出: %v\n\n", twoSum(numbers2, target2))

    // 测试用例3
    numbers3 := []int{-1, 0}
    target3 := -1
    fmt.Printf("示例3: 输入: numbers = %v, target = %d\n", numbers3, target3)
    fmt.Printf("输出: %v\n", twoSum(numbers3, target3))
}