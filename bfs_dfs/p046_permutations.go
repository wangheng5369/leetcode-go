package bfsdfs

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
	n := len(nums)
	ans := make([][]int, 0)
	used := make([]bool, n)
	path := make([]int, 0)
	var dfs func()
	dfs = func() {
		if len(path) == n {
			tmp := make([]int, n)
			copy(tmp, path)
			ans = append(ans, tmp)
		}
		for i, num := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, num)
			dfs()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs()
	return ans
}
