# 102. Binary Tree Level Order Traversal

## 題目簡述

給定一棵二元樹，回傳每一層的節點值。

```text
    3
   / \
  9   20
     /  \
    15   7
```

輸出：

```text
[[3], [9,20], [15,7]]
```

## BFS 解法

Level order 最直覺使用 BFS queue。

每一輪：

1. 先記錄 `levelSize := len(queue)`。
2. 只處理 `levelSize` 個節點。
3. 把這些節點的值放入同一個 `level`。
4. 把左右小孩放入 queue，留到下一輪處理。

## BFS 最終作答

```go
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

	return res
}
```

## DFS 版

DFS 不是天然一層一層處理，但可以用 `depth` 把節點放到正確層。

```go
func levelOrderDFS(root *TreeNode) [][]int {
	res := [][]int{}

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if len(res) == depth {
			res = append(res, []int{})
		}

		res[depth] = append(res[depth], node.Val)

		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	dfs(root, 0)
	return res
}
```

## 解法重點

- BFS 的 `levelSize` 要在每一層開始前先固定。
- 目前節點值放進 `level`。
- 小孩節點放進 `queue`，不能直接放進目前 `level`。
- DFS 版靠 `depth` 判斷該放到 `res` 的第幾層。

## 複雜度

BFS：

```text
Time: O(n)
Space: O(n)
```

DFS：

```text
Time: O(n)
Space: O(h) recursion stack，結果本身 O(n)
```

## 面試口述版

我會用 BFS queue 一層一層處理。每一輪開始先記錄目前 queue 長度，這代表目前這一層有幾個節點。接著只處理這些節點，把它們的值收集成一個 level，並把它們的左右小孩放進 queue，留到下一輪處理。每輪結束後把 level 加進結果。

如果用 DFS，也可以多帶一個 depth，讓每個節點依照深度放到對應的 `res[depth]`。
