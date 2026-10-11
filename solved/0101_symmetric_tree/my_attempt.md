# 0101. Symmetric Tree - 我的作答

## 作答時間

2026-05-10

## 完整程式碼

```go
package main

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

func isMirror(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
}

func main() {
}
```

## 你的思路

- 這題和 Same Tree 類似，但要做鏡像比較。
- 空樹是對稱的。
- 從 root 的左子樹和右子樹開始比較。
- 比較時不是左對左、右對右，而是交叉比較：
  - 左子樹的左邊 對 右子樹的右邊
  - 左子樹的右邊 對 右子樹的左邊

## 做得好的地方

- 能從 Same Tree 推導出交叉比較。
- base cases 正確：

```go
if p == nil && q == nil {
	return true
}
if p == nil || q == nil {
	return false
}
```

- 有先檢查值是否相同：

```go
if p.Val != q.Val {
	return false
}
```

- 最後遞迴比較方向正確：

```go
return isMirror(p.Left, q.Right) && isMirror(p.Right, q.Left)
```

## 可以改進的地方

目前解法已正確。這題和 Same Tree 的最大差別要記清楚：

- Same Tree：`p.Left` 對 `q.Left`，`p.Right` 對 `q.Right`
- Symmetric Tree：`left.Left` 對 `right.Right`，`left.Right` 對 `right.Left`
