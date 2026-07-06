package heap

// p347 前K个高频元素
// 给你一个整数数组 nums 和一个整数 k，请你返回其中出现频率前 k 高的元素。
// 你可以按任意顺序返回答案。
//
// 示例 1：
// 输入：nums = [1,1,1,2,2,3], k = 2
// 输出：[1,2]
//
// 示例 2：
// 输入：nums = [1], k = 1
// 输出：[1]
//
// 提示：
// - 1 <= nums.length <= 10^5
// - -10^4 <= nums[i] <= 10^4
// - 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	buckets := make([][]int, len(nums)+1)
	for k, v := range m {
		buckets[v] = append(buckets[v], k)
	}
	ans := []int{}
	for i := len(nums); i >= 1 && len(ans) < k; i-- {
		ans = append(ans, buckets[i]...)
	}
	return ans
}
