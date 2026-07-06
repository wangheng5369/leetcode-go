package binarytree

// p104 二叉树的最大深度
// 给定一个二叉树根节点 root，返回它的最大深度。
// 最大深度是从根节点到最远叶子节点的最长路径上的节点数。
//
// 提示：
// - 二叉树节点数目范围是 [0, 10^4]
// - -100 <= Node.val <= 100

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepth := 0
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		maxDepth++
	}
	return maxDepth
}
