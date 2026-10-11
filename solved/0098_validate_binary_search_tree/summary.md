# 98. Validate Binary Search Tree

## 題目簡述

給定一棵二元樹，判斷它是否為合法的 Binary Search Tree。

合法 BST 條件：

- 左子樹所有節點都必須小於目前節點。
- 右子樹所有節點都必須大於目前節點。
- 左右子樹本身也都必須是合法 BST。
- 不允許重複值。

## 最終解法

使用 DFS recursion，對每個節點帶著它允許的範圍：

```text
lower < node.Val < upper
```

初始 root 沒有上下界。往左走時，目前節點值變成新的上界；往右走時，目前節點值變成新的下界。

## 最終作答

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}

func helper(node *TreeNode, lower *int, upper *int) bool {
	if node == nil {
		return true
	}

	if lower != nil && node.Val <= *lower {
		return false
	}
	if upper != nil && node.Val >= *upper {
		return false
	}

	return helper(node.Left, lower, &node.Val) &&
		helper(node.Right, &node.Val, upper)
}
```

## 解法重點

- 不能只檢查直接父子關係。
- 每個節點都要滿足祖先節點傳下來的範圍限制。
- `nil` 節點代表沒有違規，回傳 `true`。
- 因為 BST 不允許重複，所以邊界比較要用 `<=` 和 `>=` 判斷不合法。

## 複雜度

```text
Time: O(n)
Space: O(h)
```

- `n` 是節點數，每個節點最多檢查一次。
- `h` 是樹高，來自 recursion stack。

## 面試口述版

我會用 DFS 遞迴檢查每個節點是否落在合法範圍內。初始 root 沒有上下界。往左子樹走時，目前節點值會變成新的上界；往右子樹走時，目前節點值會變成新的下界。只要某個節點不符合 `lower < val < upper`，就回傳 false。全部節點都符合則回傳 true。

## 常見反例

```text
    5
   / \
  1   4
     / \
    3   6
```

`4` 位於 `5` 的右子樹卻小於 `5`，因此整棵樹不是 BST；`3` 也違反祖先 `5` 的限制。


## 2026-09-28 面試複習

- 本輪完整作答保存在 `my_attempt.md`。既有 `final_solution.go` 是等價的 `*int` 界線版本，與上方最終作答對應；本輪使用 `*TreeNode` 傳遞祖先界線。
- 解法推導：只比較父子會漏掉更早祖先的限制。根節點界線為空；每次檢查 `lower < node.Val < upper`，往左把目前節點作新上界，往右作新下界。若有違規即回傳 false；空節點回 true。
- 正確性理由（教學整理，尚非完整獨立口述）：每次遞迴傳入的開區間恰好包含所有祖先對該位置施加的限制；更新一側界線後仍保留另一側限制。若每個節點都通過，整樹符合 BST 定義。
- 複雜度：`n` 為節點數，`h` 為樹高；每節點檢查一次，時間 `O(n)`。額外空間來自同時存在的遞迴呼叫，為 `O(h)`；平衡樹 `O(log n)`，鏈狀樹最差 `O(n)`。
- 驗證：`go run main.go` 實際輸出 `true`、`false`。暫存 Go 測試對空樹、重複值、深層祖先違規、合法深層節點及整數極值五例均通過。
- 本輪能說明：界線沿遞迴傳遞；空節點是停止條件。待複習：時間 `O(n)` 與遞迴堆疊 `O(h)` 的獨立分析，以及用樹高理解平衡與鏈狀情況。
- 下次複習：2026-09-29。
