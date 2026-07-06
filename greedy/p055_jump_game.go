package greedy

// p55 跳跃游戏
// 给定一个非负整数数组 nums，你最初位于数组的第一个索引处。
// 数组中每个元素表示你在该位置能够跳跃的最大长度。
// 判断你是否能够跳转到最后一个索引。
//
// 示例 1：
// 输入：nums = [2, 3, 1, 1, 4]
// 输出：true
// 解释：你可以跳转到最后一个索引。
// 步骤：1 -> 2 -> 3 -> 4
//
// 提示：
// - 1 <= nums.length <= 3 * 10^4
// - 0 <= nums[i] <= 10^5

func canJump(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return true
	}
	next := 0
	for i := 0; i < n-1; i++ {
		next = max(next, nums[i]+i)
		if next >= n-1 {
			return true
		}
	}
	return false
}
