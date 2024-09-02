package main

import (
	"fmt"
)

// https://leetcode.cn/problems/longest-common-prefix
func longestCommonPrefix(strs []string) string {
	// 如果字符串数组为空，直接返回空字符串
	if len(strs) == 0 {
		return ""
	}

	// 将第一个字符串假设为最长公共前缀
	prefix := strs[0]

	// 从第二个字符串开始逐一比较
	for i := 1; i < len(strs); i++ {
		// 比较当前前缀和下一个字符串，逐步缩短前缀
		for len(strs[i]) < len(prefix) || strs[i][:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1] // 每次减少前缀长度
			if len(prefix) == 0 {
				return "" // 如果前缀变为空，返回空字符串
			}
		}
	}
	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs)) // 输出 "fl"
}
