package binarytree

// p226 翻转二叉树
// 给你一个二叉树根节点 root，翻转这棵二叉树，并返回其根节点。
//
// 示例 1：
// 输入：root = [4, 2, 7, 1, 3, 6, 9]
// 输出：[4, 7, 2, 9, 6, 3, 1]
//
// 提示：
// - 二叉树节点数目范围是 [0, 100]
// - -100 <= Node.val <= 100

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = right, left
	return root
}

func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return root
}
