package binarysearch

// p81 搜索旋转排序数组 II
// 已知存在一个长度为 n 的数组，其元素按照升序排列。数组在某个 pivot 上进行了旋转，
// 使得元素变为 [nums[pivot], nums[pivot+1], ..., nums[n-1], nums[0], nums[1], ..., nums[pivot-1]] 。
//
// 例如，数组 [0, 1, 2, 4, 5, 6, 7] 在 pivot = 4 处旋转后变为 [4, 5, 6, 7, 0, 1, 2] 。
//
// 给你旋转后的数组 nums 和一个目标值 target 。如果 nums 中存在目标值 target，则返回 true，否则返回 false 。
//
// 与 p33 不同，本题 nums 中可能包含重复元素。
//
// 示例 1：
// 输入：nums = [2, 5, 6, 0, 0, 1, 2], target = 0
// 输出：true
//
// 示例 2：
// 输入：nums = [2, 5, 6, 0, 0, 1, 2], target = 3
// 输出：false
//
// 提示：
// - 1 <= nums.length <= 5000
// - -10^4 <= nums[i] <= 10^4
// - nums 可能包含重复元素
// - nums 肯定在一个点上进行了旋转
// - -10^4 <= target <= 10^4

func searchII(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}
		if nums[left] == nums[mid] {
			left++
		} else if nums[left] < nums[mid] {
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
	return false
}
