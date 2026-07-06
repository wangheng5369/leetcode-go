package trie

// p208 实现 Trie（前缀树）
// Trie（发音为 "try"）是一个树形数据结构，
// 用于高效地存储和检索字符串数据集中的键。
// 请你实现 Trie 类。
//
// Trie 类支持：
// - insert(word) 插入字符串 word
// - search(word) 如果字符串存在于 Trie 中，返回 true
// - startsWith(prefix) 如果之前插入的字符串之一具有前缀 prefix，返回 true
//
// 示例：
// 输入：
// ["Trie", "insert", "search", "search", "startsWith", "insert", "startsWith"]
// [[], ["hello"], ["hell"], ["hello"], ["hell"], ["hello"], ["hell"]]
// 输出：[null, null, false, true, true, null, true]
//
// 提示：
// - 1 <= word.length, prefix.length <= 2000
// - 最多调用 insert、search 和 startsWith 10^5 次

type Trie struct {
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
}

func (this *Trie) Search(word string) bool {
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	return false
}
