package greedy

import "sort"

// p435 无重叠区间
// 给定一个区间的数组 intervals，其中区间为 [start, end]。
// 你需要移除最少的区间，使得剩余区间不重叠。
// 请你返回需要移除的最少区间数量。
//
// 示例 1：
// 输入：intervals = [[1, 2], [2, 3], [3, 4], [1, 3]]
// 输出：1
// 解释：移除 [1, 3] 后，剩余区间都不重叠。
//
// 提示：
// - 1 <= intervals.length <= 10^5
// - intervals[i].length == 2
// - -5 * 10^4 <= start < end <= 5 * 10^4

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n < 2 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	end, ans := intervals[0][1], 0
	for i := 1; i < n; i++ {
		if intervals[i][0] < end {
			ans++
		} else {
			end = intervals[i][1]
		}
	}
	return ans
}
