# 105. Construct Binary Tree from Preorder and Inorder Traversal

## 題目簡述

給定二元樹的 `preorder` 與 `inorder` 走訪結果，重建原本的二元樹。

```text
preorder = root -> left -> right
inorder  = left -> root -> right
```

## 最終解法

使用 recursion。

每一層：

1. `preorder[0]` 是目前子樹 root。
2. 在 `inorder` 找到 root 的位置 `i`。
3. `i` 是左子樹大小。
4. 用 `i` 切出左右子樹的 preorder / inorder。
5. 遞迴建立左右子樹。

## 最終作答

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	for i, num := range inorder {
		if num != preorder[0] {
			continue
		}

		return &TreeNode{
			Val:   preorder[0],
			Left:  buildTree(preorder[1:i+1], inorder[:i]),
			Right: buildTree(preorder[1+i:], inorder[i+1:]),
		}
	}

	return nil
}
```

## Slice 切法記憶

假設 root 在 `inorder` 的位置是 `i`：

```text
左子樹大小 = i
```

所以：

```go
leftPreorder := preorder[1 : i+1]
rightPreorder := preorder[i+1:]

leftInorder := inorder[:i]
rightInorder := inorder[i+1:]
```

核心記憶：

```text
inorder 找到 root 的位置 i
i 就是左子樹節點數
preorder 的 root 後面 i 個，就是左子樹
```

## 複雜度

直覺 slice 版：

```text
Time: O(n^2)
Space: O(n)
```

- 每層都可能掃一次 `inorder` 找 root。
- recursion stack 最壞可到 `O(n)`。

## Map 優化方向

先建立：

```go
indexMap := map[int]int{}
for i, val := range inorder {
	indexMap[val] = i
}
```

之後找 root 在 `inorder` 的位置可以從 `O(n)` 變成 `O(1)`。

map 優化版通常改用 index 範圍遞迴：

```text
build(preLeft, preRight, inLeft, inRight)
```

這樣可以避免每次切 slice，也能直接使用原始 `inorder` 的 index。

優化後：

```text
Time: O(n)
Space: O(n)
```

## 面試口述版

preorder 的第一個值一定是目前子樹的 root。我先拿這個 root 去 inorder 裡找位置，因為 inorder 是 left-root-right，所以 root 左邊就是左子樹，右邊就是右子樹。root 左邊有幾個節點，就代表左子樹大小。再用這個大小去切 preorder，遞迴建立左右子樹。
