# 0110. Balanced Binary Tree - 我的作答

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

func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := height(root.Left)
	rh := height(root.Right)
	if lh == -1 || rh == -1 {
		return -1
	}
	if (lh-rh) > 1 || (rh-lh) > 1 {
		return -1
	}

	if lh > rh {
		return lh + 1
	}
	return rh + 1
}

func main() {
}
```

## 你的思路

- 這題和 Maximum Depth 類似，需要計算左右子樹高度。
- 每個節點都要檢查左右高度差是否超過 `1`。
- 使用 helper `height` 回傳高度。
- 如果某棵子樹不平衡，就用 `-1` 往上傳遞。
- 最後 `isBalanced` 只要判斷 `height(root) != -1`。

## 做得好的地方

- 使用 bottom-up recursion，一次走訪就能判斷平衡。
- `root == nil` 時高度回傳 `0`，base case 正確。
- 能把不平衡狀態用 `-1` 往上傳。
- 避免重複呼叫 `isBalanced` 導致重複計算高度。

## 批改過程中的問題

一開始 `isBalanced` 直接讀取：

```go
root.Left
root.Right
```

如果 `root == nil` 會 panic。

後來改成：

```go
return height(root) != -1
```

讓 `height` 統一處理空樹與不平衡狀態，這是更乾淨的寫法。

## 可以改進的地方

目前解法已正確。可以把變數命名從 `lh` / `rh` 改成 `leftHeight` / `rightHeight`，讓面試口述更清楚。
