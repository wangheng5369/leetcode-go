package greedy

import "sort"

// p452 用最少数量的箭引爆气球
// 给定一个二维数组 points，其中 points[i] = [start, end] 表示一个气球。
// 水平方向上射出一支箭，如果气球被射中（start <= 射箭位置 <= end），气球会被引爆。
// 请问至少需要多少支箭才能引爆所有气球？
//
// 示例 1：
// 输入：points = [[10, 16], [2, 8], [1, 6], [7, 12]]
// 输出：2
// 解释：射箭位置 x=6 可以引爆 [2, 8] 和 [7, 12]，
// 射箭位置 x=11 可以引爆 [10, 16]。
//
// 提示：
// - 1 <= points.length <= 10^5
// - points[i].length == 2
// - -2^31 <= start < end <= 2^31 - 1

func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n < 2 {
		return n
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	end := points[0][1]
	ans := 1
	for i := 1; i < n; i++ {
		if points[i][0] > end {
			ans++
			end = points[i][1]
		}
	}
	return ans
}
