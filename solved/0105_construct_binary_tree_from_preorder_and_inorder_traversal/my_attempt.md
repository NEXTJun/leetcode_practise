# 105. Construct Binary Tree from Preorder and Inorder Traversal - My Attempt

## 原始作答

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
	if len(inorder) == 0 {
		return nil
	}

	for i, num := range inorder {
		if num != preorder[0] {
			continue
		}
		return &TreeNode{
			Val:   preorder[0],
			Left:  buildTree(preorder[1:i+1], inorder[0:i]),
			Right: buildTree(preorder[1+i:], inorder[i+1:]),
		}
	}

	return nil
}

func main() {
	// Example:
	// preorder = [3,9,20,15,7]
	// inorder  = [9,3,15,20,7]
	//
	// Output tree:
	//      3
	//     / \
	//    9   20
	//       /  \
	//      15   7
}
```

## 當時思路

- `preorder` 的第一個一定是目前子樹的 root。
- 在 `inorder` 裡找到 root 的位置 `i`。
- `inorder[:i]` 是左子樹，`inorder[i+1:]` 是右子樹。
- `i` 也代表左子樹大小，所以 `preorder` 可以用這個大小切出左右子樹。

## 優點

- 有抓到 preorder / inorder 的分工：
  - preorder 負責找 root。
  - inorder 負責切左右子樹。
- slice 切法正確：
  - 左子樹 preorder：`preorder[1:i+1]`
  - 右子樹 preorder：`preorder[1+i:]`
  - 左子樹 inorder：`inorder[:i]`
  - 右子樹 inorder：`inorder[i+1:]`
- 程式簡潔，適合先在面試中寫出可工作的版本。

## 可改進

- 每次遞迴都用 `for` 掃 `inorder` 找 root，最壞時間複雜度是 `O(n^2)`。
- 可以先建立 `map[int]int` 記錄 `inorder` 中每個值的位置，將查找 root 位置改成 `O(1)`。
- map 優化版通常會改用 index 範圍遞迴，避免 slice 與原始 index 混用。
