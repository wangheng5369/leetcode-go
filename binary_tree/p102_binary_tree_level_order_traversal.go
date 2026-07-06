package binarytree

// p102 二叉树的层序遍历
// 给你二叉树的根节点 root，请返回其节点值的层序遍历结果。
// 即逐层从左向右遍历所有节点。
//
// 示例 1：
// 输入：root = [3, 9, 20, null, null, 15, 7]
// 输出：[[3], [9, 20], [15, 7]]
//
// 提示：
// - 二叉树的节点数在 [0, 2000] 范围内
// - -1000 <= Node.val <= 1000

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		step := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			step[i] = node.Val
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, step)
	}
	return ans
}
