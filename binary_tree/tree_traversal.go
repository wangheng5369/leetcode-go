package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(*TreeNode, *[]int)
	dfs = func(root *TreeNode, res *[]int) {
		if root != nil {
			dfs(root.Left, res)
			*res = append(*res, root.Val)
			dfs(root.Right, res)
		}
	}
	dfs(root, &res)
	return res
}
