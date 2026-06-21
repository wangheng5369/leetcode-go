package bfsdfs

// p079 单词搜索
// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word，
// 判断单词是否存在于网格中。
//
// 单词必须按照字母顺序，通过相邻的单元格构成，其中"相邻"单元格是水平或垂直相邻的单元格。
// 同一个单元格内的字母不能被重复使用。
//
// 示例 1：
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// 输出：true
//
// 示例 2：
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
// 输出：true
//
// 示例 3：
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
// 输出：false
//
// 提示：
// - m == board.length
// - n = board[i].length
// - 1 <= m, n <= 6
// - 1 <= word.length <= 15
// - board 和 word 仅由小写英文字母组成

func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	if n == 0 || m == 0 {
		return false
	}
	used := make([][]bool, n)
	for i := 0; i < n; i++ {
		used[i] = make([]bool, m)
	}
	var dfs func(i, j, idx int) bool
	dfs = func(i, j, idx int) bool {
		if idx == len(word) {
			return true
		}
		if i < 0 || i >= n || j < 0 || j >= m {
			return false
		}
		if used[i][j] {
			return false
		}
		if board[i][j] != word[idx] {
			return false
		}
		used[i][j] = true
		if dfs(i, j+1, idx+1) || dfs(i, j-1, idx+1) || dfs(i+1, j, idx+1) || dfs(i-1, j, idx+1) {
			return true
		}
		used[i][j] = false
		return false
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}
