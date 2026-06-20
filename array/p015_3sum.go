package array

import (
	"sort"
)

// p015 三数之和
// 给你一个整数数组 nums，判断是否存在三个元素 a, b, c 使得 a + b + c = 0？
// 找出所有和为 0 的不重复三元组。
//
// 示例 1：
// 输入：nums = [-1, 0, 1, 2, -1, -4]
// 输出：[[-1, -1, 2], [-1, 0, 1]]
//
// 示例 2：
// 输入：nums = [0, 1, 1]
// 输出：[]
//
// 提示：
// - 3 <= nums.length <= 3000
// - -10^5 <= nums[i] <= 10^5

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if sum > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return ans
}
