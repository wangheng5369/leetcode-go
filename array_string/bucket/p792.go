package bucket

type WordState struct {
	word  string
	next  int
	count int
}

type SubsequenceMatcher struct {
	wordsMap map[string]int
}

func NewMatcher(words []string) *SubsequenceMatcher {
	wMap := make(map[string]int)
	for _, word := range words {
		if len(word) > 0 {
			wMap[word]++
		}
	}
	return &SubsequenceMatcher{wordsMap: wMap}
}

func (matcher *SubsequenceMatcher) Match(s string) int {
	buckets := make(map[rune][]*WordState)
	for word, count := range matcher.wordsMap {
		firstChar := []rune(word)[0]
		buckets[firstChar] = append(buckets[firstChar], &WordState{word: word, next: 0, count: count})
	}

	total := 0
	for _, char := range s {
		waitingNodes := buckets[char]
		if len(waitingNodes) == 0 {
			continue
		}
		delete(buckets, char)

		for _, node := range waitingNodes {
			runes := []rune(node.word)
			node.next++

			if node.next == len(runes) {
				total += node.count
			} else {
				nextChar := runes[node.next]
				buckets[nextChar] = append(buckets[nextChar], node)
			}
		}
	}
	return total
}

func numMatchingSubseq(s string, words []string) int {
	matcher := NewMatcher(words)
	return matcher.Match(s)
}
