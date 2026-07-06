package dynamicprogram

// p070 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。每次你可以爬 1 或 2 个台阶。
// 你有多少种不同的方法可以爬到楼顶？
//
// 示例 1：
// 输入：n = 2
// 输出：2
//
// 提示：
// - 1 <= n <= 45

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	p, q := 2, 3
	for i := 4; i <= n; i++ {
		q += p
		p = q - p
	}
	return q
}
