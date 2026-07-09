package greedy

import "sort"

// p56 合并区间
// 给定一个区间的数组 intervals，合并所有重叠的区间。
//
// 示例 1：
// 输入：intervals = [[1, 3], [2, 6], [8, 10], [15, 18]]
// 输出：[[1, 6], [8, 10], [15, 18]]
// 解释：区间 [1, 3] 和 [2, 6] 重叠，合并为 [1, 6]。
//
// 示例 2：
// 输入：intervals = [[1, 4], [4, 5]]
// 输出：[[1, 5]]
// 解释：区间 [1, 4] 和 [4, 5] 可被视为重叠区间。
//
// 提示：
// - 1 <= intervals.length <= 10^4
// - intervals[i].length == 2
// - -5 * 10^4 <= start <= end <= 5 * 10^4

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n < 2 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{intervals[0]}
	for i := 1; i < n; i++ {
		interval := ans[len(ans)-1]
		if intervals[i][0] < interval[1] {
			interval[1] = max(interval[1], intervals[i][1])
		} else {
			ans = append(ans, intervals[i])
		}
	}
	return ans
}
