package dfs

// Given an m x n grid of characters board and a string word, return true if word exists in the grid.

// The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

// Example 1:
// Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// Output: true

func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	directions := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	var dfs func(x, y, index int) bool
	dfs = func(x, y, index int) bool {
		if x < 0 || x >= n || y < 0 || y >= m {
			return false
		}
		if board[x][y] != word[index] {
			return false
		}

		if visited[x][y] {
			return false
		}

		if index == len(word)-1 {
			return true
		}

		visited[x][y] = true
		for _, direction := range directions {
			nx, ny := direction[0]+x, direction[1]+y
			if dfs(nx, ny, index+1) {
				return true
			}
		}
		visited[x][y] = false
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == word[0] {
				return dfs(i, j, 0)
			}
		}
	}
	return false
}

func exist1(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	if n == 0 || m == 0 {
		return false
	}
	var dfs func(x, y, index int) bool
	dfs = func(x, y, index int) bool {
		if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		tmp := board[x][y]
		board[x][y] = '#'
		found := dfs(x+1, y, index+1) || dfs(x-1, y, index+1) || dfs(x, y-1, index+1) || dfs(x, y+1, index+1)
		board[x][y] = tmp
		return found
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
