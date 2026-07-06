package binarysearch

// p704 二分查找
// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target，
// 写一个函数搜索 nums 中的 target，如果存在返回其索引，否则返回 -1。
//
// 示例 1：
// 输入：nums = [-1, 0, 3, 5, 9, 12], target = 9
// 输出：4
//
// 提示：
// - 你可以假设 nums 中每个元素的取值范围都在 [-9999, 9999] 之间
// - 数组中元素唯一
// - 数组为升序排列

func search704(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
