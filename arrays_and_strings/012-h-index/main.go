package main

import (
    "fmt"
    "sort"
)

func hIndex(citations []int) int {
    // 先对数组进行排序
    sort.Ints(citations)
    
    n := len(citations)
    // 从大到小遍历排序后的数组
    for i := 0; i < n; i++ {
        if citations[i] >= n-i {
            return n-i
        }
    }
    
    return 0
}

func main() {
    citations := []int{3, 0, 6, 1, 5}
    fmt.Println(hIndex(citations)) // 输出 3
}