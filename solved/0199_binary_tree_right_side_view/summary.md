# 199. Binary Tree Right Side View

## 題目簡述

給定一棵二元樹，想像從右邊看這棵樹，回傳每一層能看到的節點值。

例子：

```text
    1
   / \
  2   3
   \   \
    5   4
```

輸出：

```text
[1,3,4]
```

注意：右視圖不是只沿著 `.Right` 走，而是每一層最右邊的節點。

## 最終解法

使用 BFS queue 按層處理。

每一層開始先固定：

```go
levelSize := len(queue)
```

這代表目前這一層有幾個節點。處理這一層時，如果目前 index 是 `levelSize-1`，代表這是這層最後一個節點，也就是右視圖看到的節點。

## 最終作答

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

	res := []int{}
	queue := []*TreeNode{root}

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

	return res
}
```

## 解法重點

- 右視圖 = 每一層最右邊的節點。
- BFS 每層開始要先固定 `levelSize`。
- `i == levelSize-1` 時加入答案。
- 左右小孩都要放進 queue，不能只放右小孩。

## 複雜度

```text
Time: O(n)
Space: O(n)
```

- `n` 是節點數，每個節點最多進出 queue 一次。
- queue 最壞可能存一層的所有節點，因此空間是 `O(n)`。

## 面試口述版

我會用 BFS 一層一層走。每一層開始先記錄 queue 長度 `levelSize`，代表這層有幾個節點。處理這層時，如果目前 index 是 `levelSize-1`，代表它是這層最後一個節點，也就是從右邊看到的節點，把它加入答案。左右小孩仍然正常放進 queue，因為右視圖不是只走右子樹。
