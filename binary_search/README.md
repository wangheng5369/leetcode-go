# 二分查找详解

---

## 一、二分查找核心思想

### 1.1 什么是二分查找

二分查找是一种在**有序数组**中查找目标元素的高效算法，每次将搜索范围缩小一半。

| 指标 | 值 |
|------|-----|
| 时间复杂度 | O(log n) |
| 空间复杂度 | O(1) |
| 前提条件 | 有序数组 |

### 1.2 二分查找的本质

```
有序数组: [1, 3, 5, 7, 9, 11, 13]
目标: 7

第1次: left=0, right=6, mid=3 → 7 < 11 → 搜索左半
第2次: left=0, right=2, mid=1 → 7 > 3 → 搜索右半
第3次: left=2, right=2, mid=2 → 7 == 7 → 找到!
```

---

## 二、二分查找通用模板

### 2.1 标准模板（左闭右闭区间）

```go
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1  // [left, right] 闭区间

    for left <= right {  // 区间不为空
        mid := left + (right-left)/2

        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1  // 目标在右半
        } else {
            right = mid - 1  // 目标在左半
        }
    }
    return -1  // 未找到
}
```

### 2.2 左开右闭区间模板

```go
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums)  // [left, right)

    for left < right {  // 区间不为空
        mid := left + (right-left)/2

        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return -1
}
```

### 2.3 模板选择规则

| 区间类型 | left 初始值 | right 初始值 | 循环条件 | 移动方向 |
|----------|-----------|-------------|----------|----------|
| [left, right] | 0 | len(nums)-1 | left <= right | left = mid+1 / right = mid-1 |
| [left, right) | 0 | len(nums) | left < right | left = mid+1 / right = mid |

---

## 三、二分查找题型分类

### 3.1 题型总览

| 题型 | 关键词 | 核心判断 |
|------|--------|----------|
| **标准二分** | 查找目标值、是否存在 | nums[mid] == target |
| **查找左边界** | 第一个 >= target、第一个满足条件 | nums[mid] >= target |
| **查找右边界** | 最后一个 <= target、最后一个满足条件 | nums[mid] <= target |
| **旋转数组** | 旋转、峰值、查找目标 | 判断哪半有序 |
| **矩阵搜索** | 二维矩阵、行/列有序 | 行列选择 |

---

## 四、标准二分

### 4.1 问题特征

| 特征 | 描述 |
|------|------|
| 关键词 | 查找、是否存在、找到 |
| 数组 | 一维有序数组 |
| 返回值 | 目标值索引 或 -1 |

### 4.2 解题模板

```go
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
```

### 4.3 经典题目

| 题目 | 标题 | 难度 |
|------|------|------|
| p704 | 二分查找 | Easy |
| p374 | 猜数字大小 | Easy |

#### p704 二分查找

```go
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
```

---

## 五、查找左边界

### 5.1 问题特征

| 特征 | 描述 |
|------|------|
| 关键词 | 第一个 >= target、第一个满足条件 |
| 返回值 | 最左满足条件的索引 |

### 5.2 解题模板

```go
func leftBound(nums []int, target int) int {
    left, right := 0, len(nums)  // [left, right)

    for left < right {
        mid := left + (right-left)/2
        if nums[mid] >= target {  // 满足条件
            right = mid           // 收缩右边界
        } else {
            left = mid + 1
        }
    }
    return left  // left 就是左边界
}
```

### 5.3 变形题目

| 题目 | 标题 |
|------|------|
| p35 | 搜索插入位置 |
| p278 | 第一个错误的版本 |

#### p35 搜索插入位置

```go
func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums)

    for left < right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}
```

---

## 六、查找右边界

### 6.1 问题特征

| 特征 | 描述 |
|------|------|
| 关键词 | 最后一个 <= target、最后一个小于等于 |
| 返回值 | 最右满足条件的索引 + 1 |

### 6.2 解题模板

```go
func rightBound(nums []int, target int) int {
    left, right := 0, len(nums)

    for left < right {
        mid := left + (right-left)/2
        if nums[mid] <= target {  // 满足条件
            left = mid + 1     // 收缩左边界
        } else {
            right = mid
        }
    }
    return left - 1  // left-1 是右边界
}
```

### 6.3 经典题目

| 题目 | 标题 | 难度 |
|------|------|------|
| p34 | 在排序数组中查找元素的第一个和最后一个位置 | Medium |

#### p34 第一个和最后一个位置

```go
func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }

    // 找左边界
    left := 0
    right := len(nums)
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            right = mid
        } else {
            left = mid + 1
        }
    }
    leftBound := left

    // 找右边界
    left = 0
    right = len(nums)
    for left < right {
        mid := left + (right-left)/2
        if nums[mid] <= target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    rightBound := left - 1

    if leftBound > rightBound {
        return []int{-1, -1}
    }
    return []int{leftBound, rightBound}
}
```

---

## 七、旋转数组

### 7.1 问题特征

| 特征 | 描述 |
|------|------|
| 数组 | 原本有序，旋转后分割成两段 |
| 关键词 | 旋转、最小值、搜索目标 |
| 核心 | 判断哪半数组是有序的 |

### 7.2 旋转数组模板

```go
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            return mid
        }

        // 判断左半是否有序
        if nums[left] <= nums[mid] {
            // 左半有序
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            // 右半有序
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    return -1
}
```

### 7.3 旋转数组系列

| 题目 | 标题 | 问法 |
|------|------|------|
| p33 | 搜索旋转排序数组 | 搜索目标值 |
| p81 | 搜索旋转排序数组 II | 搜索目标值（含重复） |
| p153 | 寻找旋转排序数组中的最小值 | 找最小值 |
| p154 | 寻找旋转排序数组中的最小值 II | 找最小值（含重复） |

#### p33 搜索旋转排序数组

```go
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            return mid
        }

        if nums[left] <= nums[mid] {  // 左半有序
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {  // 右半有序
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    return -1
}
```

#### p153 寻找旋转排序数组中的最小值

```go
func findMin(nums []int) int {
    left, right := 0, len(nums)-1

    for left < right {
        mid := left + (right-left)/2
        if nums[mid] > nums[right] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return nums[left]
}
```

---

## 八、二维矩阵搜索

### 8.1 问题特征

| 特征 | 描述 |
|------|------|
| 输入 | 二维矩阵 |
| 矩阵特点 | 每行/每列分别有序 |
| 关键词 | 搜索二维数组 |

### 8.2 解题思路

| 方法 | 适用场景 |
|------|----------|
| 右上角 | 从右上角出发，左边小，右边大 |
| 左下角 | 从左下角出发，上边小，下边大 |
| 二分查找 | 每行/每列有序时 |

#### p240 搜索二维矩阵 II

```go
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }

    // 从右上角开始
    row, col := 0, len(matrix[0])-1

    for row < len(matrix) && col >= 0 {
        if matrix[row][col] == target {
            return true
        } else if matrix[row][col] > target {
            col--  // 左移
        } else {
            row++  // 下移
        }
    }
    return false
}
```

### 8.3 经典题目

| 题目 | 标题 | 难度 |
|------|------|------|
| p74 | 搜索二维矩阵 | Medium |
| p240 | 搜索二维矩阵 II | Medium |

---

## 九、完全平方数

### 9.1 问题特征

| 特征 | 描述 |
|------|------|
| 关键词 | 完全平方数、平方根 |
| 本质 | 二分查找求 sqrt |

### 9.2 解题模板

```go
func isPerfectSquare(num int) bool {
    left, right := 0, num

    for left <= right {
        mid := left + (right-left)/2
        if mid*mid == num {
            return true
        } else if mid*mid < num {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return false
}
```

### 9.3 经典题目

| 题目 | 标题 | 难度 |
|------|------|------|
| p367 | 有效的完全平方数 | Easy |
| p69 | x 的平方根 | Easy |

---

## 十、面试高频问题

### Q1: 二分查找的边界如何确定？

| 区间类型 | left | right | 循环条件 | right 移动 |
|----------|------|-------|----------|------------|
| [left, right] | 0 | len-1 | left <= right | right = mid - 1 |
| [left, right) | 0 | len | left < right | right = mid |

### Q2: 如何避免死循环？

```go
// ❌ 错误：left = mid
left = mid      // 可能死循环
right = mid     // 可能死循环

// ✅ 正确
left = mid + 1
right = mid - 1

// 或使用 left < right 条件
```

### Q3: mid 如何计算防止溢出？

```go
// ❌ 可能溢出
mid := (left + right) / 2

// ✅ 正确
mid := left + (right-left)/2
```

### Q4: 二分查找一定能找到吗？

不一定。二分查找要求：
1. **有序数组**（或满足二分性）
2. **存在唯一解** 或 **存在性判断**

### Q5: 什么情况下用左边界，什么情况下用右边界？

| 场景 | 使用 |
|------|------|
| 找目标值存在 | 标准二分 |
| 找第一个 >= target | 左边界 |
| 找最后一个小于等于 target | 右边界 |
| 找插入位置 | 左边界 |

---

## 十一、经典题目速查

| 题目 | 题型 | 难度 | 关键点 |
|------|------|------|--------|
| p704 | 标准二分 | Easy | 基础模板 |
| p35 | 左边界 | Easy | 找插入位置 |
| p34 | 双边界 | Medium | 左边界+右边界 |
| p33 | 旋转数组 | Medium | 判断哪半有序 |
| p81 | 旋转数组 II | Medium | 含重复 |
| p153 | 旋转数组最小值 | Medium | 比较 nums[mid] 和 nums[right] |
| p154 | 旋转数组最小值 II | Hard | 含重复 |
| p74 | 二维矩阵搜索 | Medium | 右上角起点 |
| p240 | 二维矩阵搜索 II | Medium | Z 字形搜索 |
| p367 | 完全平方数 | Easy | 二分求平方根 |
| p278 | 第一个错误版本 | Easy | 左边界 |

---

## 十二、二分查找思维流程

```
开始
  │
  ├─ 数组有序？
  │    ├─ 是 → 继续
  │    └─ 否 → 找二分性（单调性）
  │
  ├─ 是标准二分？（查找目标值）
  │    └─ 是 → 标准模板
  │
  ├─ 是找左边界？（第一个 >= target）
  │    └─ 是 → 左边界模板
  │
  ├─ 是找右边界？（最后 <= target）
  │    └─ 是 → 右边界模板
  │
  ├─ 是旋转数组？
  │    └─ 是 → 判断哪半有序
  │
  └─ 是二维矩阵？
       └─ 是 → 右上角/左下角
```
