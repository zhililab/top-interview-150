package main

import "fmt"

// canCompleteCircuit 计算是否可以完成环形路线 https://leetcode.cn/problems/gas-station/
func canCompleteCircuit(gas []int, cost []int) int {
    // total_tank 记录总的油量和消耗的差值
    // current_tank 记录当前路径的油量和消耗的差值
    total_tank, current_tank := 0, 0
    // starting_station 记录当前可能的起始加油站索引
    starting_station := 0

    // 遍历所有加油站
    for i := 0; i < len(gas); i++ {
        // 更新总油量和总消耗的差值
        total_tank += gas[i] - cost[i]
        // 更新当前路径的油量和消耗的差值
        current_tank += gas[i] - cost[i]

        // 如果当前路径油量不足以到达下一个加油站
        if current_tank < 0 {
            // 将起始加油站更新为下一个加油站
            starting_station = i + 1
            // 重置当前路径油量
            current_tank = 0
        }
    }

    // 检查总油量是否足够完成环形路线
    if total_tank >= 0 {
        // 返回起始加油站索引
        return starting_station
    } else {
        // 无法完成环形路线，返回 -1
        return -1
    }
}

func main() {
    gas := []int{1, 2, 3, 4, 5}
    cost := []int{3, 4, 5, 1, 2}
    fmt.Println(canCompleteCircuit(gas, cost)) // 输出 3
}