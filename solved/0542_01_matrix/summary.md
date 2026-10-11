# 0542. 01 Matrix

## 題目

給定只包含 `0` 和 `1` 的矩陣 `mat`，回傳同樣大小的矩陣，其中每個格子的值是它到最近 `0` 的距離。每次只能上下左右移動一格。

## 最終作答

```go
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

## 解法重點

- 使用 multi-source BFS。
- 所有 `0` 都是 BFS 起點，距離為 `0`。
- 所有 `1` 先標記為 `-1`，表示還沒拜訪。
- BFS 一層一層擴散，第一次到達某格就是最近距離。
- 每次從 queue 取出一格，檢查上下左右。
- 如果鄰居沒越界且 `dist == -1`，就設定距離並加入 queue。

## 從作答改進的地方

- 原始修正版已正確。
- 批改中特別修正了 `queue := queue[1:]` 的 shadowing 問題。
- 最終版在 BFS 迴圈中加空行，讓取出節點、檢查鄰居、更新距離三段更清楚。

## 複雜度

- 時間複雜度：`O(m*n)`
- 空間複雜度：`O(m*n)`

每個格子最多進 queue 一次，每次處理最多檢查四個方向。

## 面試口述版

我會用 multi-source BFS。與其從每個 `1` 出發找最近的 `0`，不如把所有 `0` 一開始都放進 queue，距離設為 `0`，然後同時往外擴散。對於每個尚未拜訪的 `1`，第一次被 BFS 走到時，就是它離最近 `0` 的最短距離。因此我用 `dist == -1` 表示還沒拜訪；當某個鄰居第一次被走到，就設定 `dist[neighbor] = dist[current] + 1`，並放進 queue。每個格子最多處理一次，所以時間複雜度是 `O(m*n)`，空間複雜度也是 `O(m*n)`。
