# 112. Path Sum

## 題目簡述

給定一棵二元樹 `root` 和整數 `targetSum`，判斷是否存在一條從 root 到 leaf 的路徑，使路徑上的節點值總和等於 `targetSum`。

注意：必須是 root 到 leaf 的完整路徑，中途加總等於 `targetSum` 不算成功。

## 最終解法

使用 DFS recursion。

每走到一個節點，就把目前節點值從 `targetSum` 扣掉。當走到 leaf 時，檢查 leaf 值是否等於剩餘的 `targetSum`。

## 最終作答

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	remaining := targetSum - root.Val
	return hasPathSum(root.Left, remaining) ||
		hasPathSum(root.Right, remaining)
}
```

## 解法重點

- `nil` 節點沒有路徑，回傳 `false`。
- 只有 leaf 節點可以判斷是否成功。
- 遞迴往下傳的是剩餘目標值：`targetSum - root.Val`。
- 因為題目問「是否存在」，左右子樹用 `||`。

## 複雜度

```text
Time: O(n)
Space: O(h)
```

- `n` 是節點數，每個節點最多檢查一次。
- `h` 是樹高，來自 recursion stack。

## 面試口述版

我會用 DFS 從 root 往下走，每經過一個節點，就把該節點值從 targetSum 扣掉。當走到 leaf 時，檢查 leaf 的值是否等於剩下的 targetSum。如果左子樹或右子樹有任一條路徑符合，就回傳 true；如果走到空節點，代表沒有路徑，回傳 false。

## 常見反例

```text
targetSum = 5

    5
   /
  1
```

不能在 root 值等於 5 時直接成功，因為 root 不是 leaf。完整 root-to-leaf 路徑是 `5 -> 1`，總和是 6，所以答案是 false。
