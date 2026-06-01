package slidewindow

func minWindow(s string, t string) string {
	if len(s) < len(t) || t == "" {
		return ""
	}

	charCntT := make(map[byte]int)
	for i := range t {
		charCntT[t[i]]++
	}

	charCntS := make(map[byte]int)
	left := 0
	result := ""
	required := len(charCntT)
	formed := 0

	for right := 0; right < len(s); right++ {
		c := s[right]
		charCntS[c]++
		
		if charCntT[c] > 0 && charCntS[c] == charCntT[c] {
			formed++
		}

		for left <= right && formed == required {
			if result == "" || right-left+1 < len(result) {
				result = s[left : right+1]
			}

			c = s[left]
			charCntS[c]--
			if charCntT[c] > 0 && charCntS[c] < charCntT[c] {
				formed--
			}
			left++
		}
	}

	return result
}