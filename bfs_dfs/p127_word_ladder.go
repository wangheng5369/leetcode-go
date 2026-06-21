package bfsdfs

// p127 单词接龙
// 字典 wordList 中从单词 beginWord 和 endWord 的转换序列是一个按下述规则形成的序列：
// - 序列中第一个单词是 beginWord
// - 序列中最后一个单词是 endWord
// - 序列中相邻单词的相同位置仅有一个字符不同
// - 转换过程中每个单词必须存在于字典中
//
// 给定两个单词 beginWord 和 endWord 以及字典 wordList，
// 返回从 beginWord 到 endWord 的最短转换序列的单词数目。
// 如果不存在这样的转换序列，返回 0。
//
// 示例 1：
// 输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
// 输出：5
// 解释：一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog" ，返回长度为 5
//
// 示例 2：
// 输入：beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
// 输出：0
//
// 提示：
// - 1 <= beginWord.length <= 10
// - endWord.length == beginWord.length
// - 1 <= wordList.length <= 5000
// - wordList[i].length == beginWord.length
// - 字典中所有单词长度相同
// - 仅包含小写英文字母
// - 字典中可能包含重复单词
// - 可以假设 beginWord 不在字典中

func ladderLength(beginWord string, endWord string, wordList []string) int {

}
