package main

import (
	"fmt"
)

// 罗马数字到整数的映射表
var romanToIntMap = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	n := len(s)
	result := 0

	for i := 0; i < n; i++ {
		value := romanToIntMap[s[i]]

		if i < n-1 && value < romanToIntMap[s[i+1]] {
			result -= value
		} else {
			result += value
		}
	}

	return result
}

func main() {
	fmt.Println(romanToInt("III"))     // 输出: 3
	fmt.Println(romanToInt("IV"))      // 输出: 4
	fmt.Println(romanToInt("IX"))      // 输出: 9
	fmt.Println(romanToInt("LVIII"))   // 输出: 58
	fmt.Println(romanToInt("MCMXCIV")) // 输出: 1994
}
