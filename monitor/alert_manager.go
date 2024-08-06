package main

import (
	"fmt"
	"sort"
	"time"
)

// 告警结构体
type Alert struct {
	ID      string
	Source  string
	Content string
	Time    time.Time
	Level   int // 告警级别，数字越小优先级越高
	Weight  int // 告警权重，数字越大优先级越高
}

// 告警管理系统
type AlertManager struct {
	alerts []Alert
}

// 初始化告警管理系统
func NewAlertManager() *AlertManager {
	return &AlertManager{}
}

// 添加告警
func (am *AlertManager) AddAlert(alert Alert) {
	am.alerts = append(am.alerts, alert)
}

// 优先处理告警
func (am *AlertManager) ProcessAlerts() {
	// 按照权重排序告警，权重高的告警优先处理
	sort.SliceStable(am.alerts, func(i, j int) bool {
		return am.alerts[i].Weight > am.alerts[j].Weight
	})

	for _, alert := range am.alerts {
		fmt.Printf("Processing alert: %s from %s with weight %d\n", alert.Content, alert.Source, alert.Weight)
	}
}

func main() {
	am := NewAlertManager()

	// 添加告警
	alerts := []Alert{
		{ID: "1", Source: "ServiceA", Content: "CPU usage high", Time: time.Now(), Level: 1, Weight: 5},
		{ID: "2", Source: "ServiceB", Content: "Memory usage high", Time: time.Now(), Level: 2, Weight: 3},
		{ID: "3", Source: "ServiceC", Content: "Disk space low", Time: time.Now(), Level: 3, Weight: 4},
		{ID: "4", Source: "ServiceD", Content: "Network latency", Time: time.Now(), Level: 1, Weight: 2},
	}

	for _, alert := range alerts {
		am.AddAlert(alert)
	}

	// 优先处理告警
	am.ProcessAlerts()
}

/*
关键思想和数据结构：

1. `map` 数据结构：
   - `alerts` 是一个 map，用于存储所有的告警信息，键是告警的 ID，值是告警结构体 Alert。
   - `processedAlert` 是一个 map，用于跟踪已处理过的告警，键是告警的 ID，值是空结构体 struct{}，表示这个 ID 已经被处理过。

2. 去重处理：
   - 在添加告警时，通过检查 `processedAlert` map 是否存在该 ID 来实现去重。如果存在，则忽略该告警。

3. 随机数生成：
   - 使用 `rand.Intn(100)` 生成随机数作为告警的 ID，通过 `rand.Seed(time.Now().UnixNano())` 设置随机种子，以确保每次运行时生成不同的随机数序列。

4. 结构体和方法：
   - `Alert` 结构体表示一个告警信息。
   - `AlertManager` 结构体包含告警管理的逻辑，并通过方法 `AddAlert` 和 `ProcessAlerts` 来处理告警。

Go 语言语法拓展：

1. `map` 和 `struct`：
   - `map` 是 Go 语言中的哈希表，用于键值对的存储。
   - `struct` 是一种聚合数据类型，可以包含多个字段。

2. 方法和函数：
   - 方法是具有接收者的函数，定义形式为 `func (receiver ReceiverType) MethodName() ReturnType`。

3. 初始化和实例化：
   - 使用 `make(map[KeyType]ValueType)` 初始化 map。
   - 使用 `NewAlertManager()` 函数创建并返回一个 `AlertManager` 实例。

4. 随机数生成：
   - `rand.Seed(time.Now().UnixNano())` 设置随机种子。
   - `rand.Intn(n)` 生成一个介于 0 到 n-1 之间的随机整数。

*/
