package main

import (
    "fmt"
    "sort"
    "time"
)

// 告警结构体
type Alert struct {
    ID       string
    Source   string
    Content  string
    Time     time.Time
    Level    int // 告警级别，数字越小优先级越高
    Weight   int // 告警权重，数字越大优先级越高
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