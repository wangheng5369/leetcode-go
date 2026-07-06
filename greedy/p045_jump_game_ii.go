package greedy

// p45 跳跃游戏 II
// 给定一个非负整数数组 nums，你最初位于数组的第一个索引处。
// 数组中每个元素表示你在该位置能够跳跃的最大长度。
// 请你计算最少需要跳跃多少次才能跳转到最后一个索引。
// 假设你总是能够跳转到最后一个索引。
//
// 示例 1：
// 输入：nums = [2, 3, 1, 1, 4]
// 输出：2
// 解释：跳到下标 1 的位置，然后跳到最后一个索引。
//
// 提示：
// - 1 <= nums.length <= 10^4
// - 0 <= nums[i] <= 10^5

func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}
	jumps, end, next := 0, 0, 0
	for i := 0; i < n-1; i++ {
		next = max(next, nums[i]+i)
		if i == end {
			jumps++
			end = next
		}

	}
	return jumps
}
