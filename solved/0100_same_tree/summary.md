# 0100. Same Tree

## 題目

給定兩棵 binary tree `p` 和 `q`，判斷它們是否完全相同。完全相同代表結構相同，且每個對應節點的值也相同。

## 最終作答

```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
```

## 解法重點

- 使用 recursion。
- 兩邊都是 `nil`，代表兩棵空子樹相同。
- 只有一邊是 `nil`，代表結構不同。
- 兩邊節點值不同，代表樹不同。
- 目前節點相同時，遞迴比較左子樹與右子樹。
- 左子樹和右子樹都相同，整棵樹才相同。

## 從作答改進的地方

- 原始修正版邏輯正確。
- 最終版將單邊 `nil` 判斷合併成 `p == nil || q == nil`，更簡潔。
- 移除多餘括號，讓 return 更清楚。

## 複雜度

- 時間複雜度：`O(n)`
- 空間複雜度：`O(h)`

其中 `n` 是節點數，`h` 是樹高。空間來自 recursion call stack。

## 面試口述版

我會用 recursion 同時比較兩棵樹的對應節點。如果兩個節點都是 `nil`，代表這兩棵子樹都空，回傳 `true`。如果只有其中一個是 `nil`，代表結構不同，回傳 `false`。如果兩個節點值不同，也回傳 `false`。若目前節點值相同，就繼續遞迴比較左子樹和右子樹，兩邊都相同才回傳 `true`。每個節點最多比較一次，所以時間複雜度是 `O(n)`，遞迴空間是 `O(h)`。
