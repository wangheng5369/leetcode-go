package array

// p042 接雨水
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子下雨后能接多少雨水。
//
// 示例 1：
// 输入：height = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
// 输出：6
//
// 提示：
// - n == height.length
// - 1 <= n <= 2 * 10^4
// - 0 <= height[i] <= 10^5

func trap(height []int) int {
	leftMax, rightMax := height[0], height[len(height)-1]
	water := 0
	for left, right := 0, len(height)-1; left < right; {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}
	return water
}
