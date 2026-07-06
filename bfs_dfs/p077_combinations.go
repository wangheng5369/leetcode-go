package bfsdfs

// p77 组合
// 给定两个整数 n 和 k，返回所有可能的 [1, n] 范围内的 k 个数的组合。
//
// 示例 1：
// 输入：n = 4, k = 2
// 输出：
// [
//   [1,2],
//   [1,3],
//   [1,4],
//   [2,3],
//   [2,4],
//   [3,4]
// ]
//
// 示例 2：
// 输入：n = 1, k = 1
// 输出：[[1]]
//
// 提示：
// - 1 <= n <= 20
// - 1 <= k <= n

func combine(n int, k int) [][]int {
	ans := [][]int{}
	if n < k {
		return ans
	}
	path := make([]int, 0)
	used := make([]bool, n)
	var dfs func(step int)
	dfs = func(step int) {
		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for i := step; i <= n; i++ {
			if used[i-1] {
				continue
			}
			used[i-1] = true
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
			used[i-1] = false
		}
	}
	dfs(1)
	return ans
}
