package slidewindow

// p003 无重复字符的最长子串
// 给定一个字符串 s，请你找出其中不含有重复字符的最长子串的长度。
//
// 示例 1：
// 输入：s = "abcabcbb"
// 输出：3
//
// 示例 2：
// 输入：s = "bbbbb"
// 输出：1
//
// 提示：
// - 0 <= s.length <= 5 * 10^4
// - s 由英文字母、数字、符号和空格组成

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	left, right := 0, 0
	maxLen := 0
	for right < len(s) {
		if i, ok := m[s[right]]; ok && i >= left{
			left = i + 1
		}
		m[s[right]] = right
		if right-left+1 > maxLen {
			maxLen = right - left+1
		}
		right++
	}
	
	return maxLen
}
