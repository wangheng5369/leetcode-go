package array

// p283 移动零
// 给定一个数组 nums，写一个函数将数组中的所有 0 移动到数组末尾，
// 同时保持非零元素的相对顺序。
// 你能本地修改数组吗？不用返回任何东西，只需在原地修改数组即可。
//
// 示例 1：
// 输入：nums = [0, 1, 0, 3, 12]
// 输出：[1, 3, 12, 0, 0]
//
// 提示：
// - 1 <= nums.length <= 10^4
// - -2^31 <= nums[i] <= 2^31 - 1

func moveZeroes(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
}
