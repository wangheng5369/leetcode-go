package array

// p27 移除元素
// 给你一个数组 nums 和一个值 val，原地移除所有数值等于 val 的元素。
// 返回移除后数组的新长度。
// 元素的顺序可以改变。超出新长度后面的元素无论是什么都可以。
//
// 示例 1：
// 输入：nums = [3, 2, 2, 3], val = 3
// 输出：2, nums = [2, 2]
//
// 示例 2：
// 输入：nums = [0, 1, 2, 2, 3, 0, 4, 2], val = 2
// 输出：5, nums = [0, 1, 4, 0, 3]
//
// 提示：
// - 0 <= nums.length <= 300
// - 0 <= nums[i] <= 50
// - 0 <= val <= 100

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
