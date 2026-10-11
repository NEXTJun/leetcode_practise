# 0100. Same Tree - 我的作答

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

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q == nil {
		return false
	}
	if p == nil && q != nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return (isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right))
}

func main() {
}
```

## 你的思路

- 使用 recursion 比較兩棵樹。
- 如果兩邊都是 `nil`，代表兩棵空子樹相同，回傳 `true`。
- 如果只有其中一邊是 `nil`，代表結構不同，回傳 `false`。
- 如果兩邊節點值不同，回傳 `false`。
- 如果目前節點相同，繼續比較左子樹與右子樹。

## 做得好的地方

- base cases 判斷正確。
- 有先處理 `nil`，避免後面讀取 `p.Val` 或 `q.Val` 時 panic。
- 最終修正版正確比較對應子樹：

```go
isSameTree(p.Left, q.Left)
isSameTree(p.Right, q.Right)
```

## 批改過程中的問題

一開始最後遞迴比較寫成：

```go
return isSameTree(p.Left, p.Right) && isSameTree(q.Left, q.Right)
```

這是在比較同一棵樹自己的左右子樹，不是比較兩棵樹的對應位置。

正確應該是：

```go
return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
```

## 可以改進的地方

這兩段：

```go
if p != nil && q == nil {
	return false
}
if p == nil && q != nil {
	return false
}
```

可以合併成：

```go
if p == nil || q == nil {
	return false
}
```

因為前面已經處理過兩邊都 nil 的情況。
