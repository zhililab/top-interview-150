package main

import (
	"fmt"
	"strings"
)

// strStr 查找 needle 在 haystack 中第一次出现的位置
func strStr(haystack string, needle string) int {
	// 使用 strings.Index 函数查找 needle 在 haystack 中第一次出现的位置
	return strings.Index(haystack, needle)
}

func main() {
	// 测试代码
	haystack := "sadbutsad"
	needle := "sad"
	fmt.Println("Haystack:", haystack)
	fmt.Println("Needle:", needle)
	fmt.Println("First Occurrence Index:", strStr(haystack, needle))
}
