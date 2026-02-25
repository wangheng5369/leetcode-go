package string

// Given a string s, find the length of the longest substring without duplicate characters.
// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3. Note that "bca" and "cab" are also correct answers.

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	lastIndex := make([]int, 128)
	left := 0
	for right := 0; right < len(s); right++ {
		if lastIndex[s[right]] > left {
			left = lastIndex[s[right]]
		}
		lastIndex[s[right]] = right + 1
		if cur := right - left + 1; cur > maxLength {
			maxLength = cur
		}
	}
	return maxLength
}
