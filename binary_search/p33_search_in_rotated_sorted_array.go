package binarysearch

// p33 搜索旋转排序数组
// 整数数组 nums 在预先未知的某个 pivot 上进行了旋转，使得元素变为 [nums[pivot], nums[pivot+1], ..., nums[n-1], nums[0], nums[1], ..., nums[pivot-1]] 。
//
// 例如，数组 [0, 1, 2, 4, 5, 6, 7] 在 pivot = 4 处旋转后变为 [4, 5, 6, 7, 0, 1, 2] 。
//
// 给你旋转后的数组 nums 和一个目标值 target 。如果 nums 中存在目标值 target，则返回其索引，否则返回 -1 。
//
// 你必须设计时间复杂度为 O(log n) 的算法来解决此问题。
//
// 示例 1：
// 输入：nums = [4, 5, 6, 7, 0, 1, 2], target = 0
// 输出：4
//
// 示例 2：
// 输入：nums = [4, 5, 6, 7, 0, 1, 2], target = 3
// 输出：-1
//
// 示例 3：
// 输入：nums = [1], target = 0
// 输出：-1
//
// 提示：
// - 1 <= nums.length <= 5000
// - -10^4 <= nums[i] <= 10^4
// - nums 中的每个值都 独一无二
// - nums 肯定在一个点上进行了旋转
// - -10^4 <= target <= 10^4

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
