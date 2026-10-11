# 112. Path Sum - My Attempt

## 原始作答

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
		if root.Val == targetSum {
			return true
		}
	}

	t := targetSum - root.Val

	return hasPathSum(root.Left, t) || hasPathSum(root.Right, t)
}

func main() {
	// Example:
	// targetSum = 22
	//
	//        5
	//       / \
	//      4   8
	//     /   / \
	//    11  13  4
	//   /  \      \
	//  7    2      1
	//
	// Path: 5 -> 4 -> 11 -> 2
	// Output: true
}
```

## 當時思路

- 每走過一個節點，就把 `targetSum` 扣掉目前節點值。
- 空節點代表沒有路徑，回傳 `false`。
- 只有走到 leaf 時才能判斷是否成功。
- 左右子樹只要有一邊成功，整題就成功。

## 優點

- 抓到本題關鍵：成功條件必須發生在 leaf。
- 遞迴狀態轉移正確：`targetSum - root.Val`。
- 左右子樹使用 OR 合併結果，符合「存在任一路徑」的題意。

## 可改進

- leaf 判斷可以更精簡：

```go
if root.Left == nil && root.Right == nil {
	return root.Val == targetSum
}
```

- 可移除多餘空行，讓面試時寫法更乾淨。
