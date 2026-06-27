package bfsdfs

// p463 岛屿的周长
// 给定一个由 0（海洋）和 1（陆地）组成的二维网格 grid，
// 计算岛屿的周长。
//
// 示例 1：
// 输入：grid = [[0,1,0,0],[1,1,0,0],[0,1,0,0],[0,0,0,0]]
// 输出：16
//
// 示例 2：
// 输入：grid = [[1]]
// 输出：4
//
// 提示：
// - 1 <= grid.length <= 100
// - 1 <= grid[i].length <= 100
// - grid[i][j] 为 0 或 1

func islandPerimeter(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	var ans int
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] == 0 {
			ans++
			return
		}
		if grid[i][j] == 2 {
			return
		}
		grid[i][j] = 2
		dfs(i, j+1)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i-1, j)
		return
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				return ans
			}
		}
	}
	return 0
}
