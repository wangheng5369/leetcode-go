package backtracking

// p046 全排列
// 给定一个不含重复数字的数组 nums，返回其所有可能的全排列。
//
// 示例 1：
// 输入：nums = [1, 2, 3]
// 输出：[ [1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1] ]
//
// 提示：
// - 1 <= nums.length <= 8
// - -10 <= nums[i] <= 10
// - nums 中的所有整数互不相同

func permute(nums []int) [][]int {
	var ans [][]int
	used := make([]bool, len(nums))
	path := make([]int, len(nums))
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs()
	return ans
}
