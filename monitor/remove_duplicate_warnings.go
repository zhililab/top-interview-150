package main

import (
	"fmt"
)

// 去重告警
func removeDuplicateAlerts(alerts []string) []string {
	alertMap := make(map[string]bool)
	result := []string{}
	for _, alert := range alerts {
		if _, exists := alertMap[alert]; !exists {
			alertMap[alert] = true
			result = append(result, alert)
		}
	}
	return result
}

func main() {
	// 假设输入的告警信息数组
	alerts := []string{"CPU高", "内存高", "CPU高", "磁盘高", "CPU高"}
	fmt.Println(removeDuplicateAlerts(alerts))
}
