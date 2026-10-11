# 0235. Lowest Common Ancestor of a Binary Search Tree - 我的作答

## 作答時間

2026-05-16

## 完整程式碼

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func main() {
}
```

## 你的思路

- 利用 BST 性質：
  - 左子樹節點值都小於 root。
  - 右子樹節點值都大於 root。
- 如果 `p` 和 `q` 都小於 `root`，LCA 一定在左子樹。
- 如果 `p` 和 `q` 都大於 `root`，LCA 一定在右子樹。
- 否則代表 `p` 和 `q` 分在 root 兩側，或 root 本身就是其中一個節點，這時 root 就是 LCA。

## 做得好的地方

- 正確使用 BST 的大小關係。
- 沒有把它當成一般 binary tree 做全樹搜尋。
- 遞迴終止條件自然落在 `return root`。
- 可以處理 root 本身就是 `p` 或 `q` 的情況。

## 可以改進的地方

目前解法已正確。若要把空間壓到 `O(1)`，可以改用迴圈版本，不使用 recursion call stack。
