# 0235. Lowest Common Ancestor of a Binary Search Tree

## 題目

給定一棵 Binary Search Tree，以及兩個節點 `p` 和 `q`，找出它們的最近共同祖先 LCA。

BST 性質：

```text
左子樹所有值 < root.Val
右子樹所有值 > root.Val
```

## 最終作答：Recursion

```go
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
```

## 迴圈版本

```go
func lowestCommonAncestorIter(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
```

## 解法重點

- 這題是 BST，不是一般 binary tree。
- 如果 `p.Val` 和 `q.Val` 都小於 `root.Val`，答案在左子樹。
- 如果 `p.Val` 和 `q.Val` 都大於 `root.Val`，答案在右子樹。
- 否則 `root` 就是 LCA。
- 「否則」包含：
  - `p` 和 `q` 分別在 root 左右兩側。
  - `root` 本身就是 `p`。
  - `root` 本身就是 `q`。

## 從作答改進的地方

- 原始作答已正確。
- 總結中補上迴圈版本，方便說明如何把 recursion 空間從 `O(h)` 壓到 `O(1)`。

## 複雜度

Recursion：

- 時間複雜度：`O(h)`
- 空間複雜度：`O(h)`

Iteration：

- 時間複雜度：`O(h)`
- 空間複雜度：`O(1)`

其中 `h` 是樹高。

## 面試口述版

因為這是一棵 BST，所以我可以利用節點值大小決定搜尋方向。如果 `p` 和 `q` 的值都比目前 root 小，代表 LCA 一定在左子樹；如果兩者都比 root 大，代表 LCA 一定在右子樹。否則代表目前 root 剛好位在 `p` 和 `q` 的分岔點，或 root 本身就是其中一個節點，因此 root 就是最近共同祖先。每次只往一邊走，所以時間複雜度是 `O(h)`。
