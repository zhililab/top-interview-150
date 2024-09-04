package main

import (
    "fmt"
    "sort"
)

// 预处理函数，记录字符串 t 中每个字符出现的位置
func preprocess(t string) map[byte][]int {
    char_indices := make(map[byte][]int)
    for i := 0; i < len(t); i++ {
        char_indices[t[i]] = append(char_indices[t[i]], i)
    }
    return char_indices
}

// 判断字符串 s 是否为 t 的子序列
func isSubsequence(s string, t string, char_indices map[byte][]int) bool {
    prev_index := -1
    for i := 0; i < len(s); i++ {
        if indices, ok := char_indices[s[i]]; ok {
            // 使用二分查找找到大于 prev_index 的最小位置
			// 这段代码中的 func(j int) bool { return indices[j] > prev_index } 是一个闭包。
			// 闭包是一个函数，它可以捕获并包含其外部作用域中的变量。
			// 在这个例子中，闭包捕获了 prev_index 和 indices 变量。
            pos := sort.Search(len(indices), func(j int) bool { 
                return indices[j] > prev_index
            })
            if pos == len(indices) {
                return false
            }
            prev_index = indices[pos]
        } else {
            return false
        }
    }
    return true
}

func main() {
    t := "your_long_string_t"
    // 预处理字符串 t
    char_indices := preprocess(t)
    
    // 示例使用
    s1 := "example1"
    s2 := "example2"
    s3 := "your"
    s4 := "long"
    s5 := "string"
    s6 := "notasubsequence"

    // 检查 s1 和 s2 是否为 t 的子序列
    fmt.Println(isSubsequence(s1, t, char_indices)) // true or false
    fmt.Println(isSubsequence(s2, t, char_indices)) // true or false
    fmt.Println(isSubsequence(s3, t, char_indices)) // true or false
    fmt.Println(isSubsequence(s4, t, char_indices)) // true or false
    fmt.Println(isSubsequence(s5, t, char_indices)) // true or false
    fmt.Println(isSubsequence(s6, t, char_indices)) // true or false
}