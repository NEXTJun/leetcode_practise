# 0226. Invert Binary Tree

## 題目

給定一棵 binary tree，將每個節點的左右子樹交換，回傳反轉後的 root。

## 最終作答

```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
```

## 解法重點

- 使用 recursion。
- `root == nil` 是停止條件，代表空子樹。
- 對每個節點交換 `Left` 和 `Right`。
- 交換後，遞迴反轉左右子樹。
- 每個節點只需要處理一次。

## 從作答改進的地方

- 原始作答邏輯正確。
- 最終版用 Go 多重賦值交換左右子樹，程式更短：

```go
root.Left, root.Right = root.Right, root.Left
```

## 複雜度

- 時間複雜度：`O(n)`
- 空間複雜度：`O(h)`

其中 `n` 是節點數，`h` 是樹高。空間來自 recursion call stack。

## 面試口述版

我會使用 recursion。對每個節點，如果節點是 `nil`，代表空子樹，直接回傳 `nil`。否則先交換目前節點的左右子樹，接著遞迴處理新的左子樹和新的右子樹。因為每個節點只會被處理一次，所以時間複雜度是 `O(n)`。遞迴深度取決於樹高，所以空間複雜度是 `O(h)`。
