package array

// p001 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，
// 请你从数组中找出和为目标值 target 的两个整数，并返回它们的索引。
//
// 示例 1：
// 输入：nums = [2, 7, 11, 15], target = 9
// 输出：[0, 1]
// 解释：nums[0] + nums[1] = 2 + 7 = 9
//
// 提示：
// - 2 <= nums.length <= 10^4
// - -10^9 <= nums[i] <= 10^9
// - -10^9 <= target <= 10^9
// - 只会存在唯一有效答案

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{i, j}
		} else {
			m[num] = i
		}
	}
	return []int{}
}
