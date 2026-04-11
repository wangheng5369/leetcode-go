package short_url

import "strings"

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const base = 62

func Encode(id uint64) string {
	if id == 0 {
		return "0"
	}
	buf := make([]byte, 0, 8)
	for id > 0 {
		buf = append(buf, alphabet[id%base])
		id /= base
	}
	// reverse
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}

func Decode(s string) uint64 {
	var result uint64
	for _, c := range s {
		result = result*base + uint64(strings.IndexRune(alphabet, c))
	}
	return result
}
