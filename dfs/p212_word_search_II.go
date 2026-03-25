package dfs

type Trie struct {
	Root *TrieTree
}
type TrieTree struct {
	Children map[int32]*TrieTree
	IsWord   bool
}

type TrieFunc interface {
	Insert(word string)
}

func NewTrie() *Trie {
	m := map[int32]*TrieTree{}
	return &Trie{Root: &TrieTree{Children: m}}
}

func (root *Trie) Insert(word string) {
	// 1. 从 t.Root 开始
	// 2. 遍历单词 word 中的每个字符 c
	// 3. 计算 c 的索引：index := c - 'a'
	// 4. 如果 Children[index] 为 nil，创建一个新节点
	// 5. 移动到 Children[index]
	// 6. 遍历结束，标记 IsWord 为 true
	next := root.Root
	for _, idx := range word {
		if _, ok := next.Children[idx]; !ok {
			m := map[int32]*TrieTree{}
			next.Children[idx] = &TrieTree{Children: m}
		}
		next = next.Children[idx]
	}
	next.IsWord = true
}

func findWords(board [][]byte, words []string) []string {
	// 1. 初始化 Trie 并插入所有单词
	trie := NewTrie()
	for _, w := range words {
		trie.Insert(w)
	}

	res := make([]string, 0)
	rows, cols := len(board), len(board[0])

	var dfs func(r, c int, node *TrieTree, currentPath string)
	dfs = func(r, c int, node *TrieTree, currentPath string) {
		// 【你的代码：第一步】获取当前字符，检查 node.Children 中是否有这个字符
		if r < 0 || r >= rows || c < 0 || c >= cols {
			return
		}
		tmpChar := board[r][c]
		if tmpChar == '#' {
			return
		}
		if _, ok := node.Children[int32(tmpChar)]; !ok {
			return
		}
		// 如果没有，或者越界，或者已访问('#')，直接 return

		// 【你的代码：第二步】如果有，移动到子节点，更新 currentPath
		node = node.Children[int32(tmpChar)]
		currentPath += string(rune(tmpChar))
		// 【你的代码：第三步】检查子节点是否是单词结束 (IsWord)
		// 如果是，加入 res，并记得将 IsWord 设为 false 以去重
		if node.IsWord {
			res = append(res, currentPath)
			node.IsWord = false
		}
		// 【你的代码：第四步】标记当前格子为 '#'，递归搜索四个方向
		board[r][c] = '#'
		dfs(r, c+1, node, currentPath)
		dfs(r, c-1, node, currentPath)
		dfs(r+1, c, node, currentPath)
		dfs(r-1, c, node, currentPath)
		// 【你的代码：第五步】回溯：恢复当前格子
		board[r][c] = tmpChar
	}

	// 遍历网格每一个格子作为起点
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dfs(i, j, trie.Root, "")
		}
	}
	return res
}
