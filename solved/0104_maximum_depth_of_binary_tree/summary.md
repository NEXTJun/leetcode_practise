# 0104. Maximum Depth of Binary Tree

## 題目

給定一棵 binary tree，回傳它的最大深度。最大深度是從 root 到最深葉子節點的節點數。

## 最終作答

```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}
```

## 解法重點

- 使用 recursion。
- 空樹深度是 `0`。
- 一棵樹的最大深度等於：

```text
max(左子樹深度, 右子樹深度) + 1
```

- `+1` 代表目前節點本身。

## 從作答改進的地方

- 原始作答邏輯正確。
- 最終版將 `ln` / `rn` 改成 `leftDepth` / `rightDepth`，語意更清楚。
- 移除不必要的 `result` 變數，直接回傳答案。

## 複雜度

- 時間複雜度：`O(n)`
- 空間複雜度：`O(h)`

其中 `n` 是節點數，`h` 是樹高。空間來自 recursion call stack。

## 面試口述版

我會用 recursion。若目前節點是 `nil`，代表空樹，深度是 `0`。否則分別計算左子樹和右子樹的最大深度，取較大的那個再加上目前節點本身的 `1`。每個節點只會被拜訪一次，所以時間複雜度是 `O(n)`；遞迴深度取決於樹高，所以空間複雜度是 `O(h)`。
