// Given a binary string s and an integer k, return true if every binary code of length k is a substring of s. Otherwise, return false.

// Example 1:

// Input: s = "00110110", k = 2
// Output: true
// Explanation: The binary codes of length 2 are "00", "01", "10" and "11". They can be all found as substrings at indices 0, 1, 3 and 2 respectively.

func hasAllCodes(s string, k int) bool {
	codeExist := make(map[string]struct{})
	for i := 0; i < len(s)-k+1; i++ {
		code := s[i : i+k]
		if _, ok := codeExist[code]; ok {
			continue
		}
		codeExist[code] = struct{}{}
	}
	return len(codeExist) == 1<<k
}
