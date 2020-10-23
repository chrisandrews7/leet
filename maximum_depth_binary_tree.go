package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(root *TreeNode) int {
	var left int
	if root.Left != nil {
		left = 1 + max(root.Left)
	}

	var right int
	if root.Right != nil {
		right = 1 + max(root.Right)
	}

	if left > right {
		return left
	}
	return right
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(root)
}
