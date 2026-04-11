package prefixsum

// p523 连续的子数组和
// 给你一个整数数组 nums 和一个整数 k ，判断是否存在长度至少为 2 的子数组且满足其和恰好为 k 的整数倍。
// 如果存在，返回 true ；否则，返回 false 。
//
// 子数组是数组的连续部分。
//
// 示例 1：
// 输入：nums = [23, 2, 6, 4, 7], k = 6
// 输出：true
// 解释：子数组 [2, 6, 4] 的和是 12，12 是 6 的整数倍（12 % 6 = 0）。
//
// 示例 2：
// 输入：nums = [23, 2, 6, 4, 7], k = 13
// 输出：false
//
// 提示：
// - 1 <= nums.length <= 10^5
// - 0 <= nums[i] <= 10^9
// - 0 <= k <= 2^31 - 1

func checkSubarraySum(nums []int, k int) bool {
	prefixSum := 0
	prefixCount := map[int]int{0: -1} // 初始化：余数0最早出现在索引-1
	for i, num := range nums {
		prefixSum += num
		if k != 0 {
			prefixSum %= k
		}
		if j, ok := prefixCount[prefixSum]; ok {
			if i-j > 1 { // 子数组长度至少为2: i - j + 1 >= 2
				return true
			}
		} else {
			prefixCount[prefixSum] = i
		}
	}
	return false
}
