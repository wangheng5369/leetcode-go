package knapsack

// p416 分割等和子集
// 给你一个只包含正整数的非空数组 nums。
// 请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
// 示例 1：
// 输入：nums = [1, 5, 11, 5]
// 输出：true
// 解释：数组可以分割成 [1, 5, 5] 和 [11]。
//
// 示例 2：
// 输入：nums = [1, 2, 3, 5]
// 输出：false
// 解释：数组不能分割成两个和相等的子集。
//
// 提示：
// - 1 <= nums.length <= 200
// - 2 <= nums[i] <= 100

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, num := range nums {
		for i := sum; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}
	return dp[sum]
}
