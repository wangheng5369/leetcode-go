package main

import (
	"leetcode-go/array_string"
)

// FilterStream keeps only even numbers from the input slice.
func FilterStream(data []int) []int {
	res := make([]int, 0, len(data))
	for _, v := range data {
		if v%2 == 0 {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	// bigData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// result := FilterStream(bigData)
	// fmt.Println("Result:", result)
	array_string.SliceTest()
}
