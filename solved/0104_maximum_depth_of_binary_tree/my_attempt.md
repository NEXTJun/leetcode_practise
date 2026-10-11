# 0104. Maximum Depth of Binary Tree - 我的作答

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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var result int
	ln := maxDepth(root.Left)
	rn := maxDepth(root.Right)

	if ln > rn {
		result = ln + 1
	} else {
		result = rn + 1
	}
	return result
}

func main() {
}
```

## 你的思路

- 使用 recursion 計算樹高。
- 如果 `root == nil`，代表空樹，深度是 `0`。
- 分別計算左子樹深度與右子樹深度。
- 取較大的那邊再加上目前節點本身的 `1`。

## 做得好的地方

- base case 正確：

```go
if root == nil {
	return 0
}
```

- 正確使用左右子樹遞迴：

```go
ln := maxDepth(root.Left)
rn := maxDepth(root.Right)
```

- 正確回傳 `max(leftDepth, rightDepth) + 1`。

## 可以改進的地方

`result` 變數不是必要的，可以直接回傳：

```go
if ln > rn {
	return ln + 1
}
return rn + 1
```

目前解法本身已正確。
