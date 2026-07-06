package array

// p26 删除有序数组中的重复项
// 给你一个升序数组 nums，请你原地删除重复出现的元素，使每个元素只出现一次，
// 返回删除后数组的新长度。
// 不要使用额外的数组空间，必须在原地修改输入数组。
//
// 示例 1：
// 输入：nums = [1, 1, 2]
// 输出：2, nums = [1, 2]
//
// 示例 2：
// 输入：nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
// 输出：5, nums = [0, 1, 2, 3, 4]
//
// 提示：
// - 1 <= nums.length <= 3 * 10^4
// - -10^5 <= nums[i] <= 10^5

func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
