package binarysearch

// p153 寻找旋转排序数组中的最小值
// 假设一个按升序排列的数组在某个 pivot 上进行了旋转，
// 使得元素变为 [nums[pivot], nums[pivot+1], ..., nums[n-1], nums[0], nums[1], ..., nums[pivot-1]] 。
//
// 例如，数组 [0, 1, 2, 4, 5, 6, 7] 在 pivot = 4 处旋转后变为 [4, 5, 6, 7, 0, 1, 2] 。
//
// 给你旋转后的数组 nums，请你找出其中的最小元素。
// 假设数组中不存在重复元素。
//
// 你必须设计时间复杂度为 O(log n) 的算法来解决此问题。
//
// 示例 1：
// 输入：nums = [3, 4, 5, 1, 2]
// 输出：1
//
// 示例 2：
// 输入：nums = [4, 5, 6, 7, 0, 1, 2]
// 输出：0
//
// 示例 3：
// 输入：nums = [11, 13, 15, 17]
// 输出：11
//
// 提示：
// - n == nums.length
// - 1 <= n <= 5000
// - -5000 <= nums[i] <= 5000
// - nums 中的每个值都 独一无二
// - nums 肯定在一个点上进行了旋转

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		// if nums[mid] > nums[right] {
		// 	left = mid + 1
		// } else {
		// 	right = mid
		// }
		if nums[left] > nums[mid] {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return nums[left]
}
