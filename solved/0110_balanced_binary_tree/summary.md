# 0110. Balanced Binary Tree

## 題目

給定一棵 binary tree，判斷它是否高度平衡。高度平衡代表每個節點的左子樹高度與右子樹高度差不能超過 `1`。

## 最終作答

```go
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
```

## 解法重點

- 使用 bottom-up recursion。
- `height` 平常回傳子樹高度。
- 如果某棵子樹已經不平衡，`height` 回傳 `-1`。
- 當左子樹或右子樹是 `-1`，目前節點也直接回傳 `-1`。
- 若左右高度差大於 `1`，回傳 `-1`。
- 否則回傳目前子樹高度：`max(leftHeight, rightHeight) + 1`。

## 從作答改進的地方

- 原始修正版邏輯正確。
- 最終版把 `lh` / `rh` 改成 `leftHeight` / `rightHeight`，語意更明確。
- 使用 `diff` 表達高度差，讓「差距超過 1」更好口述。

## 複雜度

- 時間複雜度：`O(n)`
- 空間複雜度：`O(h)`

其中 `n` 是節點數，`h` 是樹高。空間來自 recursion call stack。

## 面試口述版

我會用 bottom-up recursion。helper `height` 回傳子樹高度，但如果發現某棵子樹已經不平衡，就回傳 `-1` 作為信號。對每個節點，我先取得左右子樹高度；如果其中一邊是 `-1`，代表下面已經不平衡，直接往上回傳 `-1`。接著檢查左右高度差是否超過 `1`，若超過也回傳 `-1`。否則回傳較大的子樹高度加一。最後 `isBalanced` 只要檢查 `height(root) != -1`。每個節點只處理一次，所以時間複雜度是 `O(n)`，空間複雜度是 `O(h)`。
