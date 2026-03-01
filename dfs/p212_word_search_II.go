package dfs

type TrieTree struct {
	Val    byte
	Next   *TrieTree
	IsWord bool
}

type Trie interface {
	Insert(word byte)
}

func (root *TrieTree) Insert(word byte) {
	root.Val = word
}

func findWords(board [][]byte, words []string) []string {

}
