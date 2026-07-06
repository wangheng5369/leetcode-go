package knapsack

// p494 目标和
// 给你一个整数数组 nums 和一个整数 target。
// 你可以在每个数字前面添加 + 或 - 符号，
// 将数组中的元素组合成表达式。
// 请你返回可以通过上述方法构造的、表达式结果等于 target 的不同表达式数目。
//
// 示例 1：
// 输入：nums = [1, 1, 1, 1, 1], target = 3
// 输出：5
// 解释：
// -1 + 1 + 1 + 1 + 1 = 3
// +1 - 1 + 1 + 1 + 1 = 3
// +1 + 1 - 1 + 1 + 1 = 3
// +1 + 1 + 1 - 1 + 1 = 3
// +1 + 1 + 1 + 1 - 1 = 3
//
// 提示：
// - 1 <= nums.length <= 20
// - 0 <= nums[i] <= 1000
// - 0 <= target <= 1000

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	p := (sum + target) / 2
	p = abs(p)
	dp := make([]int, p+1)
	dp[0] = 1
	for _, num := range nums {
		for i := p; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[p]
}

func abs(b int) int {
	if b < 0 {
		return -b
	}
	return b
}
