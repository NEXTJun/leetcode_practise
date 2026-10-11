# 0226. Invert Binary Tree - 我的作答

## 作答時間

2026-05-07

## 完整程式碼

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	rootTmp := root.Left
	root.Left = root.Right
	root.Right = rootTmp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

func main() {
}
```

## 你的思路

- 使用 recursion。
- 如果 `root == nil`，代表目前子樹是空的，直接回傳 `nil`。
- 對目前節點交換 `Left` 和 `Right`。
- 再遞迴反轉左右子樹。
- 最後回傳 `root`。

## 做得好的地方

- 選對方法：binary tree 很適合 recursion。
- 有正確停止條件：

```go
if root == nil {
	return nil
}
```

- 每個節點都會交換左右子樹。
- 交換後繼續處理子樹，邏輯正確。

## 可以改進的地方

交換左右子樹可以用 Go 多重賦值，少一個暫存變數：

```go
root.Left, root.Right = root.Right, root.Left
```

和你原本的寫法等價：

```go
rootTmp := root.Left
root.Left = root.Right
root.Right = rootTmp
```

目前解法本身已正確。
