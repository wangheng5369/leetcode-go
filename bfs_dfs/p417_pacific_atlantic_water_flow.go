package bfsdfs

// p417 太平洋大西洋水流问题
// 给定一个 m x n 的非负整数矩阵，表示一个岛屿的高度图。
// 太平洋位于矩阵的左边界和上边界，大西洋位于矩阵的右边界和下边界。
//
// 如果一个单元格的水能够流动到太平洋和大西洋，则返回这个单元格的列表。
// 水只能从高到低流动，或者从相同高度流动。
// 每个单元格只能流向上、下、左、右四个方向。
//
// 示例 1：
// 输入：heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
// 输出：[[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
//
// 提示：
// - m == heights.length
// - n == heights[i].length
// - 1 <= m, n <= 200
// - 0 <= heights[i][j] <= 10^5

func pacificAtlantic(heights [][]int) [][]int {
	n := len(heights)
	m := len(heights[0])
	if n == 0 || m == 0 {
		return [][]int{}
	}
	ans := make([][]int, 0, n)
	pacific := make([][]bool, n)
	atlantic := make([][]bool, n)
	for i := 0; i < n; i++ {
		pacific[i] = make([]bool, m)
		atlantic[i] = make([]bool, m)
	}
	var dfs func(x, y int, used [][]bool)
	dfs = func(x, y int, used [][]bool) {
		if x < 0 || x >= n || y < 0 || y >= m || used[x][y] {
			return
		}
		used[x][y] = true
		if x+1 < n && heights[x][y] <= heights[x+1][y] && !used[x+1][y] {
			dfs(x+1, y, used)
		}
		if x-1 >= 0 && heights[x][y] <= heights[x-1][y] && !used[x-1][y] {
			dfs(x-1, y, used)
		}
		if y+1 < m && heights[x][y] <= heights[x][y+1] && !used[x][y+1] {
			dfs(x, y+1, used)
		}
		if y-1 >= 0 && heights[x][y] <= heights[x][y-1] && !used[x][y-1] {
			dfs(x, y-1, used)
		}
	}

	for i := 0; i < m; i++ {
		dfs(0, i, pacific)
		dfs(n-1, i, atlantic)
	}
	for i := 0; i < n; i++ {
		dfs(i, 0, pacific)
		dfs(i, m-1, atlantic)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if pacific[i][j] && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}
