# 0542. 01 Matrix - 我的作答

## 作答時間

2026-05-16

## 完整程式碼

```go
package main

import "fmt"

func main() {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	fmt.Println(updateMatrix(mat))
}

func updateMatrix(mat [][]int) [][]int {
	rows := len(mat)
	cols := len(mat[0])

	dist := make([][]int, rows)
	for r := 0; r < rows; r++ {
		dist[r] = make([]int, cols)
	}

	queue := [][2]int{}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if mat[r][c] == 0 {
				dist[r][c] = 0
				queue = append(queue, [2]int{r, c})
			} else {
				dist[r][c] = -1
			}
		}
	}

	dirs := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		r := cell[0]
		c := cell[1]
		for _, dir := range dirs {
			nr := r + dir[0]
			nc := c + dir[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}
			if dist[nr][nc] != -1 {
				continue
			}
			dist[nr][nc] = dist[r][c] + 1
			queue = append(queue, [2]int{nr, nc})
		}
	}

	return dist
}
```

## 你的思路

- 使用 BFS。
- 從所有 `0` 同時出發，而不是從每個 `1` 去找最近的 `0`。
- `dist` 儲存每格到最近 `0` 的距離。
- `dist == -1` 表示該格還沒拜訪。
- 第一次走到某格時，就是它到最近 `0` 的最短距離。

## 做得好的地方

- 選對方法：multi-source BFS。
- 初始化時把所有 `0` 放進 queue。
- 用 `-1` 清楚標記未拜訪的 `1`。
- 使用四方向擴散。
- 距離更新正確：

```go
dist[nr][nc] = dist[r][c] + 1
```

## 批改過程中的問題

一開始寫成：

```go
queue := queue[1:]
```

這會在區塊內建立新的同名變數，造成外層 `queue` 沒有真正變短，BFS 會卡住。

正確是：

```go
queue = queue[1:]
```

這是 Go 的 shadowing 問題：

- `:=`：宣告新變數
- `=`：更新既有變數

## 為什麼不會重複計算

關鍵判斷：

```go
if dist[nr][nc] != -1 {
	continue
}
```

每個格子只有在 `dist == -1` 時才會被設定距離。一旦設定：

```go
dist[nr][nc] = dist[r][c] + 1
```

它就不再是 `-1`，後面其他路徑走到它時會直接跳過。

所以每個格子最多被成功加入 queue 一次。
