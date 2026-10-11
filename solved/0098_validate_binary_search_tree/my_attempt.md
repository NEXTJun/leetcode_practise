# 98. Validate Binary Search Tree - My Attempt

## 原始作答

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

	if lower != nil {
		if node.Val <= *lower {
			return false
		}
	}

	if upper != nil {
		if node.Val >= *upper {
			return false
		}
	}

	return helper(node.Left, lower, &node.Val) && helper(node.Right, &node.Val, upper)
}

func main() {
	// Example 1:
	//      2
	//     / \
	//    1   3
	// Output: true

	// Example 2:
	//      5
	//     / \
	//    1   4
	//       / \
	//      3   6
	// Output: false
}
```

## 當時思路

- `nil` 節點底下不會再有節點，所以可以視為合法。
- 每個節點不能只跟父節點比較，還要符合整條路徑傳下來的上下界。
- 往左子樹走時，目前節點值會變成新的 `upper`。
- 往右子樹走時，目前節點值會變成新的 `lower`。

## 優點

- 有抓到 BST 的核心：每個節點都要符合 `lower < val < upper`。
- 使用 `*int` 表示可有可無的上下界，避免用極大極小值造成邊界問題。
- 遞迴終止條件正確：`node == nil` 回傳 `true`。

## 修正過的問題

- 一開始 `isValidBST` 尚未接上 helper，後來改成 `return helper(root, nil, nil)`。
- 一開始比較沒有處理等於邊界，後來改成：
  - `node.Val <= *lower` 不合法
  - `node.Val >= *upper` 不合法
- BST 不允許重複值，所以比較必須是嚴格小於與嚴格大於。


## 2026-09-28 面試複習作答

### 原始作答（完整保留）

```go
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return check(root, nil, nil)
}

func check(node, lower, upper *TreeNode) bool {
	if node == nil {
		return true
	}
	if lower != nil && node.Val <= lower.Val {
		return false
	}
	if upper != nil && node.Val >= upper.Val {
		return false
	}
	return check(node.Left, lower, node) && check(node.Right, node, upper)
}

func main() {
	valid := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	invalid := &TreeNode{
		Val:   10,
		Left:  &TreeNode{Val: 5},
		Right: &TreeNode{Val: 15, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 20}},
	}
	fmt.Println(isValidBST(valid))   // 預期：true
	fmt.Println(isValidBST(invalid)) // 預期：false
}
```

### 本輪思路、修正與待複習

- 使用者能指出右子樹必須大於祖先，並在講解後正確給出往左更新上界的例子 (10,15) → (10,12)。
- 本輪初稿只比較父子，右子比較方向錯誤；後續初稿有左子提早回傳、界線誤用兄弟節點、未向下遞迴等問題。使用者依局部提示完成上下界遞迴。
- 使用者口述知道界線向下傳遞，並提出 `node == nil` 邊界；尚未能獨立正確分析時間與遞迴堆疊空間，曾答 O(n²)/O(1)，且將三層呼叫誤認為「三次方」。
- 執行驗證通過不等於使用者已掌握複雜度說明；此項列待複習。
