// Given a binary string s and an integer k, return true if every binary code of length k is a substring of s. Otherwise, return false.

// Example 1:

// Input: s = "00110110", k = 2
// Output: true
// Explanation: The binary codes of length 2 are "00", "01", "10" and "11". They can be all found as substrings at indices 0, 1, 3 and 2 respectively.
package main

// func hasAllCodes(s string, k int) bool {

// }

func getBinaryCodes(k int) []string {
	if k == 0 {
		return []string{""}
	}
	if k == 1 {
		return []string{"0", "1"}
	}
	arr := getBinaryCodes(k - 1)
	result := make([]string, 0, len(arr)*2)
	for _, code := range arr {
		result = append(result, code+"0", code+"1")
	}
	return result
}
