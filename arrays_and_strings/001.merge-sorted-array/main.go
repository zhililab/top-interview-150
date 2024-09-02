package main

import "fmt"

// 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
    i, j, k := m-1, n-1, m+n-1 // 初始化指针 i, j, k
    // 从后往前遍历
    for j >= 0 {
        // 如果 nums1 的元素大，则将其放入最后一个位置
        if i >= 0 && nums1[i] > nums2[j] {
            nums1[k] = nums1[i]
            i--
        } else {
            // 否则将 nums2 的元素放入最后一个位置
            nums1[k] = nums2[j]
            j--
        }
        k--
    }
}

func main() {
    nums1 := []int{1, 2, 3, 0, 0, 0} // 第一个数组，包含足够的空间
    m := 3                           // 第一个数组的有效元素数量
    nums2 := []int{2, 5, 6}          // 第二个数组
    n := 3                           // 第二个数组的元素数量

    merge(nums1, m, nums2, n)        // 调用合并函数
    fmt.Println(nums1)               // 输出合并后的数组 [1 2 2 3 5 6]
}