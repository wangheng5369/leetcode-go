package string

// p005 最长回文子串
// 给你一个字符串 s，请你找出 s 中最长的回文子串。
//
// 示例 1：
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
//
// 示例 2：
// 输入：s = "cbbd"
// 输出："bb"
//
// 提示：
// - 1 <= s.length <= 1000
// - s 仅由数字和英文字母组成

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l1, r1 := expandAroundCenter(s, i, i)
		l2, r2 := expandAroundCenter(s, i, i+1)
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start:end]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
