package main

import (
    "fmt"
    "math/rand"
    "time"
)

// RandomizedSet 结构体
type RandomizedSet struct {
    nums  []int
    index map[int]int
}

// 构造函数，初始化数据结构
func Constructor() RandomizedSet {
    rand.Seed(time.Now().UnixNano())
    return RandomizedSet{
        nums:  []int{},
        index: make(map[int]int),
    }
}

// Insert 插入一个元素
func (this *RandomizedSet) Insert(val int) bool {
    if _, exists := this.index[val]; exists {
        return false
    }
    this.index[val] = len(this.nums)
    this.nums = append(this.nums, val)
    return true
}

// Remove 删除一个元素
func (this *RandomizedSet) Remove(val int) bool {
    idx, exists := this.index[val]
    if !exists {
        return false
    }
    lastVal := this.nums[len(this.nums)-1]
    this.nums[idx] = lastVal
    this.index[lastVal] = idx
    this.nums = this.nums[:len(this.nums)-1]
    delete(this.index, val)
    return true
}

// GetRandom 获取随机元素
func (this *RandomizedSet) GetRandom() int {
    return this.nums[rand.Intn(len(this.nums))]
}

func main() {
    rs := Constructor()
    fmt.Println(rs.Insert(1))    // 输出 true
    fmt.Println(rs.Remove(2))    // 输出 false
    fmt.Println(rs.Insert(2))    // 输出 true
    fmt.Println(rs.GetRandom())  // 随机返回 1 或 2
    fmt.Println(rs.Remove(1))    // 输出 true
    fmt.Println(rs.Insert(2))    // 输出 false
    fmt.Println(rs.GetRandom())  // 随机返回 2
}