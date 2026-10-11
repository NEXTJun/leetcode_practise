package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)

	if leftHeight == -1 || rightHeight == -1 {
		return -1
	}

	diff := leftHeight - rightHeight
	if diff < 0 {
		diff = -diff
	}
	if diff > 1 {
		return -1
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
