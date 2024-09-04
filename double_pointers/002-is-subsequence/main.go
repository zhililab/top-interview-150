func isSubsequence(s string, t string) bool {
    // 1. 定义指针 i --- 指向 s 的位置 s[0]; 定义指针 j -- 指向 t 的开头位置 t[0]
    // 2. 开启遍历，如果 s[i] == t[j], 移动 i++ && j++ ，如果 j 到头，返回 false 
    // 3. 如果 s 顺利匹配完，则返回 true ,证明 s 是 t 的子序列
    i,j := 0, 0
    for i < len(s) && j < len(t) {
        if (s[i] == t[j]) {
            i++
        }
        j++
    }

    if (i == len(s)) {
        return true 
    } else {
        return false 
    }
}