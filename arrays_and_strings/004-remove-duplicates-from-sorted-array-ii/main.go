package main

import "fmt"

// 删除排序数组中的重复项 II
func removeDuplicates(nums []int) int {
    if len(nums) <= 2 {
        return len(nums)
    }

    i := 2 // 初始化慢指针，从第三个元素开始
    for j := 2; j < len(nums); j++ {
        // 如果当前元素不等于慢指针前两个元素之一，则将其放在慢指针位置
        if nums[j] != nums[i-2] {
            nums[i] = nums[j]
            i++
        }
    }
    return i // 返回数组新长度
}

func removeDuplicatesII(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    slow := 0
    for fast, v := range nums {
        if fast < 2 || nums[slow-2] != v {
            nums[slow] = v
            slow++
        }
    }
    return slow
}

func main() {
    nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
    newLength := removeDuplicates(nums)
	//newLength := removeDuplicatesII(nums)
    fmt.Println(nums[:newLength]) // 输出删除重复项后的数组
}