# 0733. Flood Fill - 我的作答

## 作答時間

2026-05-10

## 完整程式碼

```go
package main

import "fmt"

func main() {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	sr := 1
	sc := 1
	color := 2

	fmt.Println(floodFill(image, sr, sc, color))
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc]
	if originalColor == color {
		return image
	}

	var dfs func(int, int)
	dfs = func(r int, c int) {
		if r < 0 || r >= len(image) {
			return
		}
		if c < 0 || c >= len(image[0]) {
			return
		}
		if image[r][c] != originalColor {
			return
		}
		image[r][c] = color
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	dfs(sr, sc)
	return image
}
```

## 你的思路

- 使用 DFS 從起點往上下左右擴散。
- 先記住起點原本顏色 `originalColor`。
- 只處理和 `originalColor` 相同的格子。
- 超出邊界或顏色不同就停止。
- 先把目前格子改成新顏色，再繼續 DFS 四個方向。

## 做得好的地方

- 選對方法：DFS。
- 有正確記錄起點原色：

```go
originalColor := image[sr][sc]
```

- 有處理新顏色和原顏色相同的邊界：

```go
if originalColor == color {
	return image
}
```

- 先改色再遞迴，可以避免重複走回同一格。
- 邊界條件清楚。

## 可以改進的地方

目前解法已正確。若想減少重複的 DFS 呼叫語句，可以用方向陣列：

```go
dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
```

但目前手寫四個方向更直觀，面試中完全可以接受。
