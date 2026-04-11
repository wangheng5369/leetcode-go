package dfs

// p695 岛屿的最大面积
// 给你一个大小为 n x m 的二维整数网格 grid，其中：
// - 0 表示海洋
// - 1 表示陆地
//
// 每块岛屿由相邻（水平或垂直）的 1 连接而成，并且整个网格可以被海洋分隔。
//
// 请你返回网格中最大岛屿的面积。如果没有任何岛屿，则返回 0 。
//
// 相邻的定义：两个单元格共享同一条边（up, down, left, right）。
//
// 示例 1：
// 输入：grid = [
//   [0, 0, 1, 0, 0, 0, 0, 1, 0, 0],
//   [0, 0, 0, 0, 0, 0, 0, 1, 1, 1],
//   [0, 1, 1, 1, 0, 0, 0, 0, 0, 0],
//   [0, 1, 0, 0, 1, 1, 0, 0, 1, 0],
//   [0, 1, 0, 0, 1, 1, 0, 0, 0, 0],
//   [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
// ]
// 输出：6
// 解释：最大岛屿是右上角的岛屿，面积为 6。
//
// 示例 2：
// 输入：grid = [[0, 0, 0, 0, 0, 0, 0, 0]]
// 输出：0
//
// 提示：
// - n == grid.length
// - m == grid[i].length
// - 1 <= n, m <= 50
// - grid[i][j] 为 0 或 1

func maxAreaOfIsland(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	if n == 0 || m == 0 {
		return 0
	}
	res := 0
	var dfs func(x, y int) int

	dfs = func(x, y int) int {
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] != 1 {
			return 0
		}
		grid[x][y] = 0
		ans := 1
		ans += dfs(x+1, y)
		ans += dfs(x-1, y)
		ans += dfs(x, y+1)
		ans += dfs(x, y-1)
		return ans
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				res = max(res, dfs(i, j))
			}
		}
	}
	return res
}
