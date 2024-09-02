package main

import (
	"fmt"
)

// https://leetcode.cn/problems/length-of-last-word
func lengthOfLastWord(s string) int {
	length := 0
	inWord := false

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			length++
			inWord = true
		} else if inWord {
			break
		}
	}

	return length
}

func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s)) // 输出 5
}
