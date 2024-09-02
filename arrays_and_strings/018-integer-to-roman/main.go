package main

import "fmt"

// 定义罗马数字和整数值的对应关系，从大到小排列
var values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var symbols = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func intToRoman(num int) string {
	roman := ""

	// 遍历每一个罗马数字符号
	for i := 0; i < len(values); i++ {
		// 当 num 大于等于当前罗马数字值时，将对应罗马数字加入结果
		for num >= values[i] {
			roman += symbols[i]
			num -= values[i]
		}
	}

	return roman
}

func main() {
	// 测试示例
	num := 1994
	result := intToRoman(num)
	fmt.Println(result) // 输出 "MCMXCIV"
}
