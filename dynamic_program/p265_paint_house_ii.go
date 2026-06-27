package dynamicprogram

import "math"

// p265 粉刷房子 II
// 有一排房子，编号从 0 到 n-1，有 k 种颜色可以粉刷。
// 第 i 个房子粉刷成第 j 种颜色的代价为 costs[i][j]。
// 如果两个相邻的房子具有相同的颜色，那么就会触犯一些美感要求。
// 请你计算粉刷所有房子的最低成本，使得没有两个相邻的房子具有相同的颜色。
//
// 与 p256 不同的是，这里有 k 种颜色，而非固定的 3 种。
//
// 示例 1：
// 输入：n = 3, k = 3, costs = [[14, 2, 11], [11, 14, 5], [14, 3, 10]]
// 输出：10
// 解释：
// 将房子 0 涂成颜色 1（成本 = 2）
// 将房子 1 涂成颜色 0（成本 = 11）
// 将房子 2 涂成颜色 1（成本 = 10）
// 总成本 = 2 + 11 + 10 = 10
//
// 示例 2：
// 输入：n = 2, k = 2, costs = [[17, 2], [16, 16]]
// 输出：2
// 解释：
// 将房子 0 涂成颜色 1（成本 = 2）
// 将房子 1 涂成颜色 0（成本 = 16）
// 总成本 = 2 + 16 = 18（次优）
// 最优方案：房子 0 涂颜色 0（成本 17），房子 1 涂颜色 1（成本 16），总成本 = 33（次优）
// 最优：2 + 16 = 18
//
// 提示：
// - n == costs.length
// - k == costs[i].length
// - 1 <= n <= 10^5
// - 1 <= k <= 10^5
// - 1 <= costs[i][j] <= 10^4

func minCostII(costs [][]int) int {
	n, m := len(costs), len(costs[0])
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		dp[0][i] = costs[0][i]

	}
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			minCost := math.MaxInt64
			for k := 0; k < m; k++ {
				if k != j {
					minCost = min(minCost, dp[i-1][k])
				}
			}
			dp[i][j] = minCost + costs[i][j]
		}
	}
	res := math.MaxInt64
	for j := 0; j < m; j++ {
		res = min(res, dp[n-1][j])
	}
	return res
}
