package string

func maximumLength(s string) int {
	n := len(s)
	l, r := 0, n
	check := func(x int) bool {
		cnt := make([]int, 26)
		for i := 0; i < n; {
			j := i + 1
			for j < n && s[j] == s[i] {
				j++
			}
			cnt[s[i]-'a'] += max(0, j-i-x+1)
			if cnt[s[i]-'a'] >= 3 {
				return true
			}
			i = j
		}
		return false
	}
	for l < r {
		mid := (l + r + 1) >> 1
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if l == 0 {
		return -1
	}
	return l

}
