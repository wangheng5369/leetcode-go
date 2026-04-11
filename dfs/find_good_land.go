package dfs

func findGoodLand(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	if n == 0 || m == 0 {
		return 0
	}

	var dfs func(x, y int) int

	dfs = func(x, y int) int{
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
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				res = max(res, dfs(i, j))
			}
		}
	}
	return res
}
