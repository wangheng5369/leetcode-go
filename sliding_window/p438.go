package slidewindow

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	var res, temp []int
	letters := [26]int{}
	for i := 0; i < len(p); i++ {
		letters[p[i]-'a']++
		temp[s[i]-'a']++
	}

	left, right := 0, len(p)-1
	isSame := func() bool {
		for i := 0; i < 26; i++ {
			if letters[i] != temp[i] {
				return false
			}
		}
		return true
	}

	for right < len(s) {
		if isSame() {
			res = append(res, left)
		}
		temp[s[left]-'a']--
		left++
		right++
		if right < len(s) {
			temp[s[right]-'a']++
		}
	}
	return res
}
