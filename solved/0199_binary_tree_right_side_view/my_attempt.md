# 199. Binary Tree Right Side View - My Attempt

## 原始作答

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == levelSize-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	// TODO:
	// 1. root == nil ???喟征蝯?
	// 2. ??queue ??BFS
	// 3. 瘥惜?摰?levelSize := len(queue)
	// 4. ??i == levelSize-1嚗誨銵券惜?敺???暺??蝑?
	// 5. 撠酋銝璅? left ??right ?暸?queue
	return res
}

func main() {
	// Example:
	//      1
	//     / \
	//    2   3
	//     \   \
	//      5   4
	//
	// Output:
	// [1,3,4]
}
```

## 當時思路

- 右視圖不是只走 `.Right`。
- 題目要每一層從右邊看到的節點，所以可以沿用 level order BFS。
- 每層開始先固定 `levelSize := len(queue)`。
- 當 `i == levelSize-1` 時，代表目前節點是這一層最後被處理到的節點，加入答案。
- 左右小孩都要放進 queue，因為某一層即使沒有右節點，左節點也可能是右視圖看到的節點。

## 優點

- 正確使用 BFS 按層處理。
- 有固定 `levelSize`，不會把下一層混進目前層。
- 小孩加入 queue 的邏輯正確，沒有只走右子樹。
- 時間與空間複雜度都符合預期。

## 可改進

- 正式答案可刪除 scaffold 留下的亂碼 TODO。
- `queue` 可以直接初始化：

```go
queue := []*TreeNode{root}
```
