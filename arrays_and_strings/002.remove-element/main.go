package main

import "fmt"

// 移除指定元素
func removeElement(nums []int, val int) int {
    k := 0 // k表示有效元素的计数
    for _, num := range nums {
        if num != val { // 如果当前元素不等于指定值
            nums[k] = num // 将其移动到数组前部
            k++
        }
    }
    return k // 返回移除后的数组新长度
}

func main() {
    nums := []int{3, 2, 2, 3}
    val := 3

    newLength := removeElement(nums, val) // 调用移除元素函数
    fmt.Println(nums[:newLength]) // 输出移除后的数组
}