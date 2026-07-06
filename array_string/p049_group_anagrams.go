package arraystring

// p49 字母异位词分组
// 给你一个字符串数组，请你将字母异位词组合在一起。
// 字母异位词是由相同字母重排构成的字符串。
//
// 示例 1：
// 输入：strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出：[["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]
//
// 示例 2：
// 输入：strs = [""]
// 输出：[[""]]
//
// 示例 3：
// 输入：strs = ["a"]
// 输出：[["a"]]
//
// 提示：
// - 1 <= strs.length <= 10^4
// - 0 <= strs[i].length <= 100
// - strs[i] 由小写英文字母组成

func groupAnagrams(strs []string) [][]string {
	maps := make(map[[26]int][]string)
	for _, str := range strs {
		arr := [26]int{}
		for _, ch := range str {
			arr[ch-'a']++
		}
		maps[arr] = append(maps[arr], str)
	}
	ans := make([][]string, 0)
	for _, strings := range maps {
		ans = append(ans, strings)
	}
	return ans
}
