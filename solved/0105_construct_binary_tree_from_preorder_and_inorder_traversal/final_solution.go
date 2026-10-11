package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	for i, num := range inorder {
		if num != preorder[0] {
			continue
		}

		return &TreeNode{
			Val:   preorder[0],
			Left:  buildTree(preorder[1:i+1], inorder[:i]),
			Right: buildTree(preorder[1+i:], inorder[i+1:]),
		}
	}

	return nil
}
