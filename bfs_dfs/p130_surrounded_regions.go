package dfs

// p130 被围绕的区域
// 给你一个 m x n 的矩阵 board，其中：
// - 'X' 代表海洋（不可通行）
// - 'O' 代表陆地
//
// 捕获所有不被任何海洋围绕的陆地区域，即找到所有与边界相连的 'O' 区域，
// 将其余的 'O' 区域全部变成 'X'。
//
// 被围绕的定义为：不与边界上的任何 'O' 相连的 'O' 区域。
//
// 示例 1：
// 输入：board = [
//   ['X', 'X', 'X', 'X'],
//   ['X', 'O', 'O', 'X'],
//   ['X', 'X', 'O', 'X'],
//   ['X', 'O', 'X', 'X']
// ]
// 输出：[
//   ['X', 'X', 'X', 'X'],
//   ['X', 'X', 'X', 'X'],
//   ['X', 'X', 'X', 'X'],
//   ['X', 'O', 'X', 'X']
// ]
// 解释：右下角的 'O' 与边界相连，不被捕获，所以保留。
//
// 示例 2：
// 输入：board = [
//   ['X']
// ]
// 输出：[['X']]
//
// 提示：
// - m == board.length
// - n == board[i].length
// - 1 <= m, n <= 200
// - board[i][j] 为 'X' 或 'O'

func solve(board [][]byte) {
	n, m := len(board), len(board[0])
	if n == 0 || m == 0 {
		return
	}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != 'O' {
			return
		}
		board[x][y] = 'A'
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}

	for i := 0; i < n; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][m-1] == 'O' {
			dfs(i, m-1)
		}
	}
	for i := 0; i < m; i++ {
		if board[0][i] == 'O' {
			dfs(0, i)
		}
		if board[n-1][i] == 'O' {
			dfs(n-1, i)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case 'A':
				board[i][j] = 'O'
			}
		}
	}
}
