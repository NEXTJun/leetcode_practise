# 0101. Symmetric Tree

## 題目

給定一棵 binary tree，判斷它是否左右對稱。

## 最終作答

```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}
```

## 解法重點

- 使用 recursion。
- 空樹是對稱的。
- 判斷一棵樹是否對稱，本質是判斷 root 的左右子樹是否互為鏡像。
- 兩邊都是 `nil`，代表這一段對稱。
- 只有一邊是 `nil`，代表結構不對稱。
- 節點值不同，代表不對稱。
- 節點值相同時，要交叉比較：
  - `left.Left` 對 `right.Right`
  - `left.Right` 對 `right.Left`

## 從作答改進的地方

- 原始作答已正確。
- 最終版只把 helper 參數名稱從 `p` / `q` 改成 `left` / `right`，讓鏡像語意更明確。

## 複雜度

- 時間複雜度：`O(n)`
- 空間複雜度：`O(h)`

其中 `n` 是節點數，`h` 是樹高。空間來自 recursion call stack。

## 面試口述版

我會用 recursion。空樹本身是對稱的；如果 root 不為空，就比較 root 的左子樹和右子樹是否互為鏡像。鏡像比較時，如果兩邊都是 nil，回傳 true；如果只有一邊是 nil，回傳 false；如果兩邊值不同，也回傳 false。若目前兩個節點值相同，接著交叉比較：左節點的左子樹要對上右節點的右子樹，左節點的右子樹要對上右節點的左子樹。每個節點最多比較一次，所以時間複雜度是 `O(n)`，遞迴深度是樹高，所以空間複雜度是 `O(h)`。
