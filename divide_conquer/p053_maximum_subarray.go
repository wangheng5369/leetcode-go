package divideconquer

// p53 最大子数组和
// 给你一个整数数组 nums，请你在其中找出具有最大和的连续子数组，
// 返回其最大和。
//
// 子数组 是数组中的一个连续子部分。
//
// 示例 1：
// 输入：nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
// 输出：6
// 解释：连续子数组 [4, -1, 2, 1] 的和最大，为 6。
//
// 提示：
// - 1 <= nums.length <= 10^5
// - -10^5 <= nums[i] <= 10^5

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	maxSum := nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		maxSum = max(maxSum, dp[i])
	}
	return maxSum
}

func maxSubArray1(nums []int) int {
	curSum := 0
	maxSum := nums[0]
	for _, num := range nums {
		if num > 0 {
			maxSum = max(maxSum, curSum+num)
		}
		curSum += num
		if curSum < 0 {
			curSum = 0
		}
	}
	return maxSum
}

func maxSubArray2(nums []int) int {

}
