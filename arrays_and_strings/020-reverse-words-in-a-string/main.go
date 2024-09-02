package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// 将字符串分割为单词数组，并自动去除多余空格
	words := strings.Fields(s)
	/* for _, word := range words {
		fmt.Println(word)
	} */

	fmt.Println(words[0])
	fmt.Println(words[1])

	// 反转单词数组
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// 拼接单词数组为结果字符串
	return strings.Join(words, " ")
}

func main() {
	s := "  the sky  is blue  "
	fmt.Println(reverseWords(s)) // 输出 "blue is sky the"
}
