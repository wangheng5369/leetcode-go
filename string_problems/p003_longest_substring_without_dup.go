package slidewindow

// Given a string s, find the length of the longest substring without duplicate characters.
// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	left, right := 0, 0
	m := make(map[byte]int)
	for right < len(s) {
		if i, ok := m[s[right]]; ok {
			left = i + 1
		}
		m[s[right]] = right
		maxLen = max(maxLen, right-left+1)
		right++
	}
	return maxLen
}
