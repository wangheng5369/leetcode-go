# DFS/BFS 问题详解

---

## 一、DFS vs BFS 对比

| 维度 | DFS | BFS |
|------|-----|-----|
| 全称 | 深度优先搜索 | 广度优先搜索 |
| 数据结构 | 栈（递归） | 队列 |
| 搜索顺序 | 先深入后回溯 | 先扩展后深入 |
| 适用场景 | 找所有解、连通性、路径 | 最短路径、层级遍历 |
| 空间复杂度 | O(h)，h为树高 | O(2^h)，满二叉树最后一层 |

---

## 二、DFS 题型分类

### 2.1 题型总览

| 题型 | 关键词 | 模板 |
|------|--------|------|
| **岛屿系列** | 岛屿、数量、面积、周长 | 方向数组 + 递归 |
| **网格搜索** | 单词搜索、路径 | visited + 递归 |
| **连通分量** | 被围绕的区域、连接 | 边界出发/全局染色 |
| **回溯算法** | 组合、排列、子集 | 选择列表 + 回溯 |

---

## 三、岛屿系列

### 3.1 岛屿系列题目

| 题目 | 标题 | 问法 |
|------|------|------|
| p200 | 岛屿数量 | 有多少个岛屿 |
| p695 | 岛屿的最大面积 | 最大岛屿面积 |
| p463 | 岛屿的周长 | 岛屿周长是多少 |
| p130 | 被围绕的区域 | 哪些区域被围绕 |

### 3.2 岛屿通用模板

```go
// 方向数组：上下左右
var directions = [][]int{{-1,0}, {1,0}, {0,-1}, {0,1}}

func dfs(grid [][]int, i, j int) {
    // 1. 越界检查
    if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
        return
    }
    // 2. 不符合条件检查（如不是岛屿）
    if grid[i][j] != 1 {
        return
    }
    // 3. 标记已访问（避免重复访问）
    grid[i][j] = 0  // 或用 visited 数组

    // 4. 递归访问四个方向
    dfs(grid, i-1, j)
    dfs(grid, i+1, j)
    dfs(grid, i, j-1)
    dfs(grid, i, j+1)
}
```

### 3.3 p200 岛屿数量

```go
func numIslands(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    count := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' {
                dfs(grid, i, j)
                count++
            }
        }
    }
    return count
}
```

### 3.4 p695 岛屿的最大面积

```go
func maxAreaOfIsland(grid [][]int) int {
    if len(grid) == 0 {
        return 0
    }
    maxArea := 0
    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
            return 0
        }
        grid[i][j] = 0
        area := 1
        area += dfs(i-1, j)
        area += dfs(i+1, j)
        area += dfs(i, j-1)
        area += dfs(i, j+1)
        return area
    }

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                maxArea = max(maxArea, dfs(i, j))
            }
        }
    }
    return maxArea
}
```

---

## 四、网格搜索

### 4.1 单词搜索系列

| 题目 | 标题 | 问法 |
|------|------|------|
| p79 | 单词搜索 | 单词是否存在 |
| p212 | 单词搜索 II | 所有能拼出的单词 |

### 4.2 单词搜索模板

```go
func exist(board [][]byte, word string) bool {
    if len(board) == 0 || len(board[0]) == 0 {
        return false
    }
    m, n := len(board), len(board[0])

    var dfs func(i, j, idx int) bool
    dfs = func(i, j, idx int) bool {
        // 1. 找到完整单词
        if idx == len(word) {
            return true
        }
        // 2. 越界检查
        if i < 0 || i >= m || j < 0 || j >= n {
            return false
        }
        // 3. 不匹配检查
        if board[i][j] != word[idx] {
            return false
        }

        // 4. 标记已访问 + 递归
        board[i][j] = '#'  // 临时标记
        dirs := [][]int{{-1,0}, {1,0}, {0,-1}, {0,1}}
        for _, d := range dirs {
            if dfs(i+d[0], j+d[1], idx+1) {
                return true
            }
        }
        board[i][j] = word[idx]  // 撤销标记
        return false
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if dfs(i, j, 0) {
                return true
            }
        }
    }
    return false
}
```

---

## 五、连通分量（染色法）

### 5.1 核心思路

| 方法 | 适用场景 |
|------|----------|
| 边界出发染色 | 被围绕的区域（O变X） |
| 全局染色 | 太平洋大西洋水流 |

### 5.2 p130 被围绕的区域

```go
func solve(board [][]byte) {
    if len(board) == 0 {
        return
    }
    m, n := len(board), len(board[0])

    // 从边界'O'出发，标记所有与边界相连的'O'
    var dfs func(i, j int)
    dfs = func(i, j int) {
        if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
            return
        }
        board[i][j] = 'A'  // 标记为可达
        dfs(i-1, j)
        dfs(i+1, j)
        dfs(i, j-1)
        dfs(i, j+1)
    }

    // 从四条边界出发
    for i := 0; i < m; i++ {
        dfs(i, 0)
        dfs(i, n-1)
    }
    for j := 0; j < n; j++ {
        dfs(0, j)
        dfs(m-1, j)
    }

    // 二次遍历：A保持'O'，其余'O'变'X'
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'A' {
                board[i][j] = 'O'
            } else if board[i][j] == 'O' {
                board[i][j] = 'X'
            }
        }
    }
}
```

### 5.3 p417 太平洋大西洋水流

```go
func pacificAtlantic(heights [][]int) [][]int {
    if len(heights) == 0 {
        return [][]int{}
    }
    m, n := len(heights), len(heights[0])

    pacific := make([][]bool, m)
    atlantic := make([][]bool, m)
    for i := 0; i < m; i++ {
        pacific[i] = make([]bool, n)
        atlantic[i] = make([]bool, n)
    }

    var dfs func(i, j int, visited [][]bool)
    dfs = func(i, j int, visited [][]bool) {
        if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] {
            return
        }
        visited[i][j] = true
        dirs := [][]int{{-1,0}, {1,0}, {0,-1}, {0,1}}
        for _, d := range dirs {
            ni, nj := i+d[0], j+d[1]
            if ni >= 0 && ni < m && nj >= 0 && nj < n && heights[ni][nj] >= heights[i][j] {
                dfs(ni, nj, visited)
            }
        }
    }

    // 从太平洋边界(左+上)出发
    for i := 0; i < m; i++ {
        dfs(i, 0, pacific)
    }
    for j := 0; j < n; j++ {
        dfs(0, j, pacific)
    }
    // 从大西洋边界(右+下)出发
    for i := 0; i < m; i++ {
        dfs(i, n-1, atlantic)
    }
    for j := 0; j < n; j++ {
        dfs(m-1, j, atlantic)
    }

    // 取交集
    var result [][]int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if pacific[i][j] && atlantic[i][j] {
                result = append(result, []int{i, j})
            }
        }
    }
    return result
}
```

---

## 六、回溯算法

### 6.1 回溯题型分类

| 题型 | 关键词 | 例子 |
|------|--------|------|
| **子集** | 子集、组合、集合 | p78 子集 |
| **组合** | 组合、选择、和为target | p77 组合 |
| **排列** | 全排列、排列 | p46 全排列 |
| **切割** | 切割、拆分 | p131 分割回文串 |
| **棋盘** | N皇后、数独 | p51 N皇后 |

### 6.2 回溯通用模板

```go
func backtrack(路径, 选择列表) {
    if 满足结束条件 {
        result = append(result, 路径)
        return
    }
    for 选择 in 选择列表 {
        // 做选择
        路径 = append(路径, 选择)
        // 标记已访问
        used[选择] = true

        // 递归
        backtrack(路径, 选择列表)

        // 撤销选择
        路径 = 路径[:len(路径)-1]
        used[选择] = false
    }
}
```

### 6.3 子集问题模板

```go
func subsets(nums []int) [][]int {
    result := [][]int{}
    path := []int{}

    var backtrack func(start int)
    backtrack = func(start int) {
        // 子集问题：每个节点都是解
        tmp := make([]int, len(path))
        copy(tmp, path)
        result = append(result, tmp)

        for i := start; i < len(nums); i++ {
            path = append(path, nums[i])
            backtrack(i + 1)
            path = path[:len(path)-1]
        }
    }

    backtrack(0)
    return result
}
```

### 6.4 组合问题模板

```go
func combine(n, k int) [][]int {
    result := [][]int{}
    path := []int{}

    var backtrack func(start int)
    backtrack = func(start int) {
        if len(path) == k {
            tmp := make([]int, k)
            copy(tmp, path)
            result = append(result, tmp)
            return
        }
        for i := start; i <= n; i++ {
            path = append(path, i)
            backtrack(i + 1)
            path = path[:len(path)-1]
        }
    }

    backtrack(1)
    return result
}
```

### 6.5 排列问题模板

```go
func permute(nums []int) [][]int {
    result := [][]int{}
    path := []int{}
    used := make([]bool, len(nums))

    var backtrack func()
    backtrack = func() {
        if len(path) == len(nums) {
            tmp := make([]int, len(nums))
            copy(tmp, path)
            result = append(result, tmp)
            return
        }
        for i := 0; i < len(nums); i++ {
            if used[i] {
                continue
            }
            used[i] = true
            path = append(path, nums[i])
            backtrack()
            path = path[:len(path)-1]
            used[i] = false
        }
    }

    backtrack()
    return result
}
```

---

## 七、经典题目速查

| 题目 | 题型 | 难度 | 关键点 |
|------|------|------|--------|
| p200 | 岛屿数量 | Medium | 方向数组 + 递归 |
| p695 | 岛屿的最大面积 | Medium | 返回最大面积 |
| p463 | 岛屿的周长 | Easy | 岛屿周长计算 |
| p79 | 单词搜索 | Medium | visited 避免重复 |
| p130 | 被围绕的区域 | Medium | 边界染色法 |
| p417 | 太平洋大西洋水流 | Medium | 双边界染色 |
| p46 | 全排列 | Medium | 回溯 + used 数组 |
| p77 | 组合 | Medium | 回溯 + start 参数 |
| p78 | 子集 | Medium | 回溯收集所有节点 |
| p212 | 单词搜索 II | Hard | Trie + 回溯 |

---

## 八、面试高频问题

### Q1: 什么时候用 DFS，什么时候用 BFS？

| 场景 | 推荐 | 原因 |
|------|------|------|
| 找最短路径 | BFS | 层级遍历，最早遇到的就是最短 |
| 找所有解 | DFS | 可以系统地遍历所有分支 |
| 连通分量 | DFS | 递归深入，回溯简洁 |
| 层级遍历 | BFS | 队列天然按层 |

### Q2: DFS 如何避免重复访问？

| 方法 | 适用场景 |
|------|----------|
| 修改原数组值 | 岛屿问题（grid[i][j] = 0） |
| used 数组 | 排列问题 |
| 标记字符 | 单词搜索（board[i][j] = '#'） |

### Q3: 回溯和 DFS 的区别？

- DFS 是搜索策略（深度优先）
- 回溯是 DFS 的一种实现（用递归 + 撤销选择）
- 回溯 = DFS + 剪枝 + 状态恢复
