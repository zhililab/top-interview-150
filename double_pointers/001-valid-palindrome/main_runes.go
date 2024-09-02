package main

import (
	"fmt"
	"time"
	"runtime"
)

func isPalindrome(s string) bool {
	runes := []rune(s) // 将字符串转换为 rune 列表
	left, right := 0, len(runes) - 1
    
    for left < right {
        // 跳过非字母数字字符
        for left < right && !isAlphanumeric(runes[left]) {
            left++
        }
        for left < right && !isAlphanumeric(runes[right]) {
            right--
        }
        
        // 比较字符
        if toLowerCase(runes[left]) != toLowerCase(runes[right]) {
            return false
        }
        
        left++
        right--
	}
	return true
}

func isAlphanumeric(c rune) bool {
    return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func toLowerCase(c rune) rune {
    if c >= 'A' && c <= 'Z' {
        return c + 32
    }
    return c
}

func main() {
    testCases := []string{
        "A man, a plan, a canal: Panama",
        "race a car",
        " ",
        "No 'x' in Nixon",
    }

    start := time.Now() // 记录开始时间
    var memStats runtime.MemStats
    runtime.ReadMemStats(&memStats) // 记录内存使用情况

    for _, testCase := range testCases {
        result := isPalindrome(testCase)
        fmt.Printf("Is \"%s\" a palindrome? %v\n", testCase, result)
    }

    // 记录结束时间和内存使用情况
    elapsed := time.Since(start)
    runtime.ReadMemStats(&memStats)
    fmt.Printf("Time taken: %s, Memory allocated: %d bytes\n", elapsed, memStats.Alloc)
}
