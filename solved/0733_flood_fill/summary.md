# 0733. Flood Fill

## 題目

給定二維整數陣列 `image`、起點 `(sr, sc)` 和新顏色 `color`。把起點所在區塊中，所有和起點原本顏色相同、且上下左右相連的格子，都改成新顏色。

## 最終作答

```go
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

## 解法重點

- 使用 DFS 從起點擴散。
- 擴散條件不是「遇到 0 停」，而是只走和起點原顏色相同的格子。
- `originalColor == color` 時要直接回傳，避免重複遞迴。
- 每次進入有效格子後，先改成新顏色，再往上下左右繼續。
- 超出邊界或顏色不同時停止。

## 從作答改進的地方

- 原始作答已正確。
- 總結中補強關鍵觀念：停止條件是 `image[r][c] != originalColor`，不是固定遇到某個顏色停止。
- 保留手寫四方向，方便初學時看清楚 DFS 流程。

## 複雜度

- 時間複雜度：`O(m*n)`
- 空間複雜度：`O(m*n)`

最壞情況整張圖都要填色，DFS call stack 最深也可能到 `m*n`。

## 面試口述版

我會先記錄起點原本的顏色 `originalColor`。如果原顏色已經等於新顏色，直接回傳，避免重複遞迴。接著從起點開始做 DFS；每次進入一個格子時，先檢查是否超出邊界，以及目前格子是否仍是 `originalColor`。如果不是，就停止。若是有效格子，就把它改成新顏色，然後往上下左右四個方向繼續 DFS。最壞情況會拜訪整張圖，所以時間複雜度是 `O(m*n)`，遞迴空間最壞也是 `O(m*n)`。
