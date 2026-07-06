package heap

// p215 数组中的第K个最大元素
// 给定整数数组 nums 和整数 k，
// 请你找出数组中第 k 个最大的元素。
// 请注意，是第 k 个最大元素（不是第 k 个不同元素）。
//
// 示例 1：
// 输入：nums = [3, 2, 1, 5, 6, 4], k = 2
// 输出：5
//
// 解释：排序后的数组是 [1, 2, 3, 4, 5, 6]，
// 第 2 个最大的元素是 5。
//
// 提示：
// - 1 <= k <= nums.length <= 10^4
// - -10^4 <= nums[i] <= 10^4

func findKthLargest(nums []int, k int) int {
	// 桶
	bucket := make([]int, 20001)
	for i := 0; i < len(nums); i++ {
		bucket[nums[i]+10000]++
	}
	for i := 20001; i >= 0; i-- {
		k = k - bucket[i]
		if k == 0 {
			return i - 10000
		}
	}
	return -1
}
