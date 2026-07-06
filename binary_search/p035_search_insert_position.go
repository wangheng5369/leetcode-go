package binarysearch

// p35 搜索插入位置
// 给定一个排序数组和一个目标值，
// 如果目标值存在于数组中，返回其索引。
// 否则，返回按顺序插入位置。
//
// 示例 1：
// 输入：nums = [1, 3, 5, 6], target = 5
// 输出：2
//
// 示例 2：
// 输入：nums = [1, 3, 5, 6], target = 2
// 输出：1
//
// 提示：
// - 1 <= nums.length <= 10^4
// - -10^4 <= nums[i] <= 10^4
// - nums 无重复元素
// - nums 已排序

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
