package main

import (
	"fmt"
	"strings"
)

// zigzagConvert 将字符串按照指定的行数进行之字形转换
func zigzagConvert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	// 创建一个切片来存储每一行的字符串
	rows := make([]strings.Builder, numRows)
	currentRow := 0
	goingDown := false

	// 遍历字符串中的每个字符
	for _, char := range s {
		rows[currentRow].WriteRune(char)
		if currentRow == 0 || currentRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			currentRow++
		} else {
			currentRow--
		}
	}

	// 将所有行的内容合并成一个结果字符串
	var result strings.Builder
	for _, row := range rows {
		result.WriteString(row.String())
	}

	return result.String()
}

func main() {
	// 测试代码
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println("Input String:", s)
	fmt.Println("Number of Rows:", numRows)
	fmt.Println("Zigzag Conversion:", zigzagConvert(s, numRows))
}
