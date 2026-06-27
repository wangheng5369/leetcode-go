# 背包问题详解

---

## 一、问法分类总览

| 问法类型 | 关键词 | dp 定义 | dp 初始化 | 外层循环 | 内层循环 | 转移方程 |
|----------|--------|----------|-----------|----------|----------|----------|
| **能否装满** | 能不能、是否存在、分割 | `dp[j]` = 能否装满容量 j | `dp[0]=true` | 物品 | 逆序 | `dp[j] = dp[j] \|\| dp[j-num]` |
| **最值问题** | 最少、最多、最大、最小 | `dp[j]` = 容量 j 的最值 | `dp[0]=0, 其他=极值` | 物品 | 正序/逆序 | `dp[j] = min/max(dp[j], dp[j-num]+1)` |
| **组合数** | 多少种、有几种、组合数 | `dp[j]` = 凑成 j 的方法数 | `dp[0]=1` | 物品 | 正序 | `dp[j] += dp[j-num]` |
| **排列数** | 排列数、顺序有关 | `dp[j]` = 凑成 j 的排列数 | `dp[0]=1` | 容量 | 正序 | `dp[j] += dp[j-coin]` |
| **二维费用** | 两个约束、0和1的个数 | `dp[j][k]` = 两种容量下的最值 | `dp[0][0]=0` | 物品 | 逆序 | `dp[j][k] = max(dp[j][k], dp[j-num][k-1]+1)` |
| **差值问题** | 两堆差值最小 | `dp[j]` = 容量 j 能装的最大价值 | `dp[0]=0` | 物品 | 逆序 | `dp[j] = max(dp[j], dp[j-num]+num)` |

---

## 二、能否装满问题

### 2.1 问题特征

| 特征 | 描述 |
|------|------|
| 关键词 | 能不能、是否存在、是否可以、分割成 |
| 问法 | 能不能分成两个相等的子集？能不能凑出 target？ |
| 返回值 | true / false |

### 2.2 dp 设置

```go
dp := make([]bool, target+1)
dp[0] = true  // 容量为0，一定能装满（什么都不选）
```

### 2.3 循环设置

| 设置项 | 值 |
|--------|-----|
| 外层 | 遍历物品 |
| 内层 | **逆序** (target → item) |
| 原因 | 每个物品只能用一次，不能重复选择 |

### 2.4 转移方程

```go
dp[j] = dp[j] || dp[j-num]  // 选或不选
```

### 2.5 典型题目

#### p416 分割等和子集

```go
func canPartition(nums []int) bool {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    if sum%2 == 1 {  // 奇数无法平分
        return false
    }
    target := sum / 2

    dp := make([]bool, target+1)
    dp[0] = true

    for _, num := range nums {
        for j := target; j >= num; j-- {
            dp[j] = dp[j] || dp[j-num]
        }
    }
    return dp[target]
}
```

---

## 三、最值问题

### 3.1 问题特征

| 特征 | 描述 |
|------|------|
| 关键词 | 最少、最多、最大、最小 |
| 问法 | 最少需要多少硬币？最多能抢劫多少钱？ |
| 返回值 | int（最小/最大数量/价值） |

### 3.2 dp 设置

```go
dp := make([]int, target+1)
for i := range dp {
    dp[i] = 极值  // 通常是 target+1 或 -inf
}
dp[0] = 0
```

### 3.3 循环设置

| 问题类型 | 外层 | 内层 | 原因 |
|----------|------|------|------|
| 硬币最少 | 物品 | **正序** | 每种硬币可以用无限次 |
| 物品最值 | 物品 | **逆序** | 每个物品只能用一次 |

### 3.4 转移方程

```go
// 最少
dp[j] = min(dp[j], dp[j-num]+1)

// 最多
dp[j] = max(dp[j], dp[j-num]+num)
```

### 3.5 典型题目

#### p322 零钱兑换（完全背包，最少）

```go
func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := range dp {
        dp[i] = amount + 1  // 初始化为不可能的大值
    }
    dp[0] = 0

    for _, coin := range coins {
        for j := coin; j <= amount; j++ {  // 正序：硬币可以用无限次
            dp[j] = min(dp[j], dp[j-coin]+1)
        }
    }
    if dp[amount] > amount {
        return -1
    }
    return dp[amount]
}
```

#### p198 打家劫舍（0/1 背包，最多）

```go
func rob(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    dp := make([]int, n)
    dp[0] = nums[0]
    if n > 1 {
        dp[1] = max(nums[0], nums[1])
    }
    for i := 2; i < n; i++ {
        dp[i] = max(dp[i-1], dp[i-2]+nums[i])  // 逆序隐含在递推中
    }
    return dp[n-1]
}
```

---

## 四、组合数问题

### 4.1 问题特征

| 特征 | 描述 |
|------|------|
| 关键词 | 多少种、有几种、组合数、方法数 |
| 问法 | 凑成 amount 有多少种方法？有几种组合？ |
| 返回值 | int（方法数） |

### 4.2 dp 设置

```go
dp := make([]int, target+1)
dp[0] = 1  // 容量为0，只有一种方法（什么都不选）
```

### 4.3 循环设置

| 设置项 | 值 |
|--------|-----|
| 外层 | **遍历物品** |
| 内层 | **正序** (item → target) |
| 原因 | 每种物品可以用无限次；外层物品保证每个物品只被计入一次 |

### 4.4 转移方程

```go
dp[j] += dp[j-num]
```

### 4.5 典型题目

#### p518 零钱兑换 II

```go
func change(amount int, coins []int) int {
    dp := make([]int, amount+1)
    dp[0] = 1

    for _, coin := range coins {
        for j := coin; j <= amount; j++ {  // 正序
            dp[j] += dp[j-coin]
        }
    }
    return dp[amount]
}
```

---

## 五、排列数问题

### 5.1 问题特征

| 特征 | 描述 |
|------|------|
| 关键词 | 排列数、顺序、组合 |
| 问法 | 有多少种排列方式？顺序重要吗？ |
| 返回值 | int（排列数） |

### 5.2 与组合数的区别

| 类型 | 外层 | 内层 | 区别 |
|------|------|------|------|
| **组合数** | 物品 | 正序 | 物品顺序固定 |
| **排列数** | 容量 | 正序 | 先遍历容量，再遍历物品 |

### 5.3 循环设置

```go
// 排列数：外层容量，内层物品
dp := make([]int, amount+1)
dp[0] = 1
for j := 1; j <= amount; j++ {       // 外层：容量
    for _, coin := range coins {       // 内层：物品
        if j >= coin {
            dp[j] += dp[j-coin]
        }
    }
}
```

### 5.4 典型题目

#### 排列数 vs 组合数对比

```go
// 组合数（先遍历物品）p518
for _, coin := range coins {      // 外层物品
    for j := coin; j <= amount; j++ {
        dp[j] += dp[j-coin]
    }
}

// 排列数（先遍历容量）p377
for j := 1; j <= amount; j++ {   // 外层容量
    for _, num := range nums {
        if j >= num {
            dp[j] += dp[j-num]
        }
    }
}
```

#### p377 组合总和 IV

```go
func combinationSum4(nums []int, target int) int {
    dp := make([]int, target+1)
    dp[0] = 1
    for j := 1; j <= target; j++ {       // 外层容量
        for _, num := range nums {          // 内层物品
            if j >= num {
                dp[j] += dp[j-num]
            }
        }
    }
    return dp[target]
}
```

---

## 六、二维费用问题

### 6.1 问题特征

| 特征 | 描述 |
|------|------|
| 关键词 | 两个约束、0的个数和1的个数、两种费用 |
| 问法 | 0不超过m个，1不超过n个？两种资源同时限制？ |

### 6.2 dp 设置

```go
dp := make([][]int, m+1)
for i := range dp {
    dp[i] = make([]int, n+1)
}
// dp[i][j] = 消耗i个0和j个1的最大物品数
```

### 6.3 循环设置

| 设置项 | 值 |
|--------|-----|
| 外层 | 遍历物品 |
| 内层 | **逆序**（两层都逆序） |

### 6.4 转移方程

```go
dp[j][k] = max(dp[j][k], dp[j-zeros][k-ones]+1)
```

### 6.5 典型题目

#### p474 一和零

```go
func findMaxForm(strs []string, m, n int) int {
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }

    for _, s := range strs {
        zeros := strings.Count(s, "0")
        ones := len(s) - zeros
        for j := m; j >= zeros; j-- {       // 逆序
            for k := n; k >= ones; k-- {    // 逆序
                dp[j][k] = max(dp[j][k], dp[j-zeros][k-ones]+1)
            }
        }
    }
    return dp[m][n]
}
```

---

## 七、差值问题

### 7.1 问题特征

| 特征 | 描述 |
|------|------|
| 关键词 | 两堆差值最小、分成两堆 |
| 问法 | 分成两堆，差值最小是多少？ |

### 7.2 核心转化

```
设两堆石头分别为 A 和 B
A + B = sum
|A - B| = sum - 2*min(A, B)

要使差值最小 → 使 min(A, B) 最大 → 也就是使其中一堆尽可能接近 sum/2

bag = sum / 2
差值 = sum - 2*dp[bag]
```

### 7.3 循环设置

| 设置项 | 值 |
|--------|-----|
| 外层 | 遍历物品（石头） |
| 内层 | **逆序** |

### 7.4 典型题目

#### p1049 最后一块石头的重量 II

```go
func lastStoneWeightII(stones []int) int {
    sum := 0
    for _, v := range stones {
        sum += v
    }
    target := sum / 2

    dp := make([]int, target+1)
    for _, stone := range stones {
        for j := target; j >= stone; j-- {
            dp[j] = max(dp[j], dp[j-stone]+stone)
        }
    }
    return sum - 2*dp[target]
}
```

---

## 八、目标和问题（正负号组合）

### 8.1 问题特征

| 特征 | 描述 |
|------|------|
| 关键词 | 正负号、+/-、表达式 |
| 问法 | 添加 +/- 号，有多少种方式得到 target？ |

### 8.2 核心转化

```
设正号的和为 P，负号的和为 N
P + N = sum(nums)
P - N = target
-----------------
2*P = target + sum
P = (target + sum) / 2
```

### 8.3 循环设置

| 设置项 | 值 |
|--------|-----|
| 外层 | 遍历物品 |
| 内层 | **逆序** |

### 8.4 典型题目

#### p494 目标和

```go
func findTargetSumWays(nums []int, target int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    if abs(target) > sum {
        return 0
    }
    if (target+sum)%2 == 1 {
        return 0
    }
    bag := (target + sum) / 2

    dp := make([]int, bag+1)
    dp[0] = 1
    for _, num := range nums {
        for j := bag; j >= num; j-- {
            dp[j] += dp[j-num]
        }
    }
    return dp[bag]
}
```

---

## 九、问法速查表

| 问法 | dp 定义 | dp[0] | 外层 | 内层 | 转移方程 | 典型题目 |
|------|---------|-------|------|------|----------|----------|
| 能否装满 | bool | true | 物品 | 逆序 | `dp[j] \|\| dp[j-num]` | p416 |
| 最少硬币 | int | 0/极值 | 物品 | **正序** | `min(dp[j], dp[j-num]+1)` | p322 |
| 最多价值 | int | 0 | 物品 | 逆序 | `max(dp[j], dp[j-num]+num)` | p198 |
| 组合数 | int | 1 | 物品 | **正序** | `dp[j] += dp[j-num]` | p518 |
| 排列数 | int | 1 | **容量** | 正序 | `dp[j] += dp[j-coin]` | p377 |
| 二维费用 | int[][] | 0 | 物品 | 逆序(两层) | `dp[j][k] = max(...)` | p474 |
| 差值最小 | int | 0 | 物品 | 逆序 | `sum - 2*dp[target]` | p1049 |
| 目标和 | int | 1 | 物品 | 逆序 | `dp[j] += dp[j-num]` | p494 |
