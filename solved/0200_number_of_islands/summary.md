# 200. Number of Islands

## 題目簡述

給定一個由 `'1'` 和 `'0'` 組成的 grid：

```text
'1' = land
'0' = water
```

上下左右相連的 land 算同一座島。請回傳總共有幾座島。

## 最終解法

使用 DFS。

掃過每一格：

1. 遇到 `'0'` 或已拜訪過的格子就跳過。
2. 遇到 `'1'`，代表找到一座新島，`count++`。
3. 從這格開始 DFS，把上下左右相連的 `'1'` 全部標記成已拜訪。

## 最終作答

```go
package main

func numIslands(grid [][]byte) int {
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= len(grid) {
			return
		}
		if c < 0 || c >= len(grid[0]) {
			return
		}
		if grid[r][c] != '1' {
			return
		}

		grid[r][c] = '2'

		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	count := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] != '1' {
				continue
			}
			count++
			dfs(r, c)
		}
	}

	return count
}
```

## 解法重點

- 這題等於「每找到一個新的 `'1'`，就做一次 Flood Fill」。
- DFS 只處理值為 `'1'` 的格子。
- visited 可以原地標記成 `'2'` 或 `'0'`。
- `count++` 只能在發現新島的起點時做一次。

## 複雜度

```text
Time: O(m*n)
Space: O(m*n)
```

- 每個格子最多被拜訪一次。
- DFS recursion stack 最壞可能包含整張 grid。

## 面試口述版

我會掃過 grid 每一格。當遇到 `'1'`，代表找到一座新的島，所以 count 加一。接著從這格開始做 DFS，往上下左右擴散，把所有相連的 `'1'` 標記成已拜訪，例如改成 `'2'`。這樣之後掃到同一座島的其他格子時，就不會再次計算。

## 類似題

- 733 Flood Fill：從指定起點擴散。
- 200 Number of Islands：掃完整張圖，每遇到新起點就擴散一次。


## 2026-09-28 面試複習

- 本輪完整作答原文見 `my_attempt.md`。既有 `final_solution.go` 為原地標記版本，與上方最終作答對應；本輪使用 `map[[2]int]bool` 保存已訪問座標，不修改輸入。
- 正確性：外層只在遇到未訪問的陸地時加一；DFS 將從該格可達的所有陸地標記，因此同一座島的其他格不會再加一。四方向搜尋恰好符合題目連通定義，水與越界會停止。
- 複雜度：設 `m` 列、`n` 欄。外層掃描及 DFS 合計處理每格常數次，時間平均 `O(m*n)`（map 操作按平均 O(1) 計）。本輪的 visited map 最多 `O(m*n)`，DFS 堆疊最深 `O(m*n)`，合計額外空間 `O(m*n)`。既有原地標記版仍有最差 `O(m*n)` 遞迴堆疊。
- 驗證：`go run main.go` 實際輸出 `2`、`1`；暫存 Go 測試對空網格、空列、單格陸地、單格水、對角線分離、全部相連、矩形分離七例均通過。測試不等於窮舉證明。
- 本輪口述已確認：不重複計數理由、時間 O(m*n)、額外空間 O(m*n)；停止條件與邊界案例主要由實作及測試驗證，尚未取得獨立完整口述。
- 下次複習：2026-10-05。
