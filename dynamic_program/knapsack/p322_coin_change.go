package knapsack

import "math"

// p322 零钱兑换
// 给你一个整数数组 coins 表示不同面额的硬币，和一个整数 amount 表示总金额。
// 请你计算并返回凑成总金额所需的最少硬币个数。
// 如果没有任何硬币组合能凑成总金额，返回 -1。
// 每种硬币的数量无限使用。
//
// 示例 1：
// 输入：coins = [1, 2, 5], amount = 11
// 输出：3
// 解释：11 = 5 + 5 + 1
//
// 示例 2：
// 输入：coins = [2], amount = 3
// 输出：-1
//
// 示例 3：
// 输入：coins = [1], amount = 0
// 输出：0
//
// 提示：
// - 1 <= coins.length <= 12
// - 1 <= coins[i] <= 2^31 - 1
// - 0 <= amount <= 10^4

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
