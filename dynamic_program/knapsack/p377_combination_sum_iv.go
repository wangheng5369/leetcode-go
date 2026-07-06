package knapsack

// p377 组合总和 IV
// 给你一个由不同整数组成的数组 nums 和一个目标整数 target。
// 请你找出 nums 中所有可以使数字和为 target 的排列的组合数。
//
// 排列的顺序不同视为不同的组合。
//
// 示例 1：
// 输入：nums = [1, 2, 3], target = 4
// 输出：7
// 解释：
// 可能的排列组合为：
// [1, 1, 1, 1]
// [1, 1, 2]
// [1, 2, 1]
// [1, 3]
// [2, 1, 1]
// [2, 2]
// [3, 1]
// 共 7 种。
//
// 示例 2：
// 输入：nums = [9], target = 3
// 输出：0
//
// 提示：
// - 1 <= nums.length <= 200
// - 1 <= nums[i] <= 50
// - 1 <= target <= 1000

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
