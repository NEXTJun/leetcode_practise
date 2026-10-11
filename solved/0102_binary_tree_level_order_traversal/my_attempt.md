# 102. Binary Tree Level Order Traversal - My Attempt

## 原始作答

```go
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}
	// TODO:
	// 1. root == nil ???喟征蝯?
	// 2. 撱箇? queue嚗???root
	// 3. 瘥?頛芸?閮?levelSize := len(queue)
	// 4. ?? levelSize ??暺??園???撅斤???	// 5. ?椰?喳?摮拙???queue
	return res
}

func main() {
	// Example:
	//      3
	//     / \
	//    9   20
	//       /  \
	//      15   7
	//
	// Output:
	// [[3], [9,20], [15,7]]
}
```

## 當時思路

- 題目要一層一層輸出，所以用 BFS queue。
- 每一輪開始先固定 `levelSize := len(queue)`。
- 本輪只處理 `levelSize` 個節點，這些節點就是同一層。
- 目前節點的值放進 `level`。
- 左右小孩放進 `queue`，等下一層處理。

## 修正過的問題

- 一開始寫成 `queue := queue[1:]`，這會建立新的區域變數，應改成 `queue = queue[1:]`。
- 一開始把小孩的值直接放進 `level`，會把下一層混到目前層；正確做法是把小孩節點放進 `queue`。
- 最後要回傳 `res`，不是 `nil`。

## DFS 版補充

DFS 也能做 level order，重點是多帶一個 `depth`：

- `depth` 表示目前節點在哪一層。
- 第一次到某一層時，先替 `res[depth]` 建立空 slice。
- 把目前節點值 append 到 `res[depth]`。
- 遞迴走左右子樹，深度加一。
