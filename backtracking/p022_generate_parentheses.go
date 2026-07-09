package backtracking

// p22 括号生成
// 数字 n 代表生成括号的对数，请你设计一个函数，用于生成所有可能的且有效的括号组合。
//
// 示例 1：
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
//
// 示例 2：
// 输入：n = 1
// 输出：["()"]
//
// 提示：
// - 1 <= n <= 8

func generateParenthesis(n int) []string {
	ans := []string{}
	var dfs func(string, int, int)
	dfs = func(s string, left int, right int) {
		if left > n || left < right {
			return
		}
		if left == n && right == n {
			ans = append(ans, s)
			return
		}
		dfs(s+"(", left+1, right)
		dfs(s+")", left, right+1)
	}
	dfs("", 0, 0)
	return ans
}
