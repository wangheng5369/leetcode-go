package dynamicprogram

// p256 粉刷房子
// 假设你有 n 个房子，排成一排，编号从 0 到 n-1。
// 每个房子可以被粉刷成三种颜色之一：红色、蓝色或绿色。
// 第 i 个房子粉刷成红色、蓝色或绿色的代价分别为 costs[i][0]、costs[i][1] 和 costs[i][2]。
// 如果相邻的两个房子粉刷成同一种颜色，那么就会触犯一些美感要求。
// 请你计算粉刷所有房子的最低成本，使得没有两个相邻的房子具有相同的颜色。
//
// 示例 1：
// 输入：costs = [[17, 2, 17], [16, 16, 5], [14, 3, 19]]
// 输出：10
// 解释：
// 将房子 0 涂成蓝色，成本为 2
// 将房子 1 涂成绿色，成本为 5
// 将房子 2 涂成绿色，成本为 3
// 总成本 = 2 + 5 + 3 = 10
//
// 示例 2：
// 输入：costs = [[7, 6, 2]]
// 输出：2
//
// 提示：
// - costs.length == n
// - costs[i].length == 3
// - 1 <= n <= 10^5
// - 1 <= costs[i][j] <= 10^4

func minCost(costs [][]int) int {
	// 1 定义dp[i][j]表示第i号房子涂成j颜色的最小成本，其中j取0，1，2
	dp := make([][]int, len(costs))
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	// 2 初始化0号房子每种颜色是costs[0][j]
	dp[0][0] = costs[0][0]
	dp[0][1] = costs[0][1]
	dp[0][2] = costs[0][2]
	// 3 状态转换：dp[i][j] = min(dp[i-1][j1]) + costs[i][j],其中j1与j不同
	for i := 1; i < len(costs); i++ {
		// 涂成红色
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i][0]
		// 涂成蓝色
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i][1]
		// 涂成绿色
		dp[i][2] = min(dp[i-1][0], dp[i-1][1]) + costs[i][2]
	}
	return min(dp[len(costs)-1][0], min(dp[len(costs)-1][2], dp[len(costs)-1][1]))
}
