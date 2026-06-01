package prefixsum

// p560 和为 K 的子数组
// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回该数组中和为 k 的连续子数组的个数。
//
// 示例 1：
// 输入：nums = [1, 1, 1], k = 2
// 输出：2
// 解释：连续的子数组 [1, 1]（索引 0-1）和 [1, 1]（索引 1-2）之和都为 2。
//
// 示例 2：
// 输入：nums = [1, 2, 3], k = 3
// 输出：2
// 解释：子数组 [1, 2]（索引 0-1）和 [3]（索引 2-2）之和都为 3。
//
// 提示：
// - 1 <= nums.length <= 2 * 10^4
// - -1000 <= nums[i] <= 1000
// - -10^7 <= k <= 10^7


func subarraySum(nums []int, k int) int {
	prefixCount := make(map[int]int)
	prefixCount[0] = 1 // base case for subarray starting at index 0
	result := 0
	prefixSum := 0
	for _, num := range nums {
		prefixSum += num
		if count, ok := prefixCount[prefixSum-k]; ok {
			result += count
		}
		prefixCount[prefixSum]++

	}
	return result
}
