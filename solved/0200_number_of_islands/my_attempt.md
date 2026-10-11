# 200. Number of Islands - My Attempt

## 原始作答

```go
package main

func numIslands(grid [][]byte) int {

	var dfs func(int, int)
	dfs = func(r int, c int) {
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
	var count int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '1' {
				continue
			}
			count++
			dfs(i, j)
		}
	}
	// TODO:
	// 1. ??瘥???	// 2. ? '1'嚗誨銵冽?唬?摨扳撜塚?count++
	// 3. 敺?? DFS嚗?銝?撌血?賊?? '1' ?券璅??歇?赤
	// 4. ?敺???count
	return count
}

func main() {
	// Example:
	// grid = [
	//   ['1','1','0','0','0'],
	//   ['1','1','0','0','0'],
	//   ['0','0','1','0','0'],
	//   ['0','0','0','1','1'],
	// ]
	//
	// Output: 3
}
```

## 當時思路

- 這題和 733 Flood Fill 很像。
- 掃過整張 grid，遇到尚未拜訪的 `'1'`，代表找到一座新島。
- `count++` 後，用 DFS 把同一座島上下左右相連的 `'1'` 全部標記成已拜訪。
- 之後掃到同一座島的其他格子時，因為已經不是 `'1'`，不會重複計算。

## 修正過的問題

- 一開始把 grid 當數字處理，寫成 `1` / `2`。
- 這題型別是 `[][]byte`，LeetCode 給的是字元 `'1'` / `'0'`，所以要用：

```go
grid[r][c] != '1'
```

- visited 可以標記成 `'2'` 或直接改成 `'0'`；只要之後不再等於 `'1'` 就可以。

## 優點

- DFS 越界條件正確。
- 四方向擴散正確。
- `count++` 放在遇到新島時，而不是 DFS 裡每格都加。
- 使用原地標記，不需要額外 visited 陣列。


## 2026-09-28 面試複習作答

### 原始作答（完整保留）

```go
package main

import "fmt"

func numIslands(grid [][]byte) int {
	var res int
	visited := make(map[[2]int]bool)
	// TODO: 逐格掃描；遇到尚未處理的陸地時，島嶼數加一。
	var len_r, len_c int
	len_r = len(grid)
	if len_r == 0 {
		return 0
	}
	len_c = len(grid[0])
	if len_c == 0 {
		return 0
	}
	for i := 0; i < len_r; i++ {
		for j := 0; j < len_c; j++ {
			if visited[[2]int{i, j}] {
				continue
			}
			if grid[i][j] == '1' {
				res++
				markIsland(grid, i, j, visited)
			}
		}
	}
	// TODO: 從該格往上下左右擴展，並在處理每格時立即標記。
	return res
}

func markIsland(grid [][]byte, row, col int, visited map[[2]int]bool) {
	len_r := len(grid)
	len_c := len(grid[0])
	if row < 0 || row > len_r-1 || col < 0 || col > len_c-1 {
		return
	}
	if !visited[[2]int{row, col}] && grid[row][col] == '1' {
		visited[[2]int{row, col}] = true
		markIsland(grid, row-1, col, visited)
		markIsland(grid, row+1, col, visited)
		markIsland(grid, row, col-1, visited)
		markIsland(grid, row, col+1, visited)
	}
}

func main() {
	grid1 := [][]byte{[]byte("110"), []byte("100"), []byte("001")}
	grid2 := [][]byte{[]byte("111"), []byte("010"), []byte("111")}
	fmt.Println(numIslands(grid1)) // 預期：2
	fmt.Println(numIslands(grid2)) // 預期：1
}
```

### 本輪思路與修正

- 使用者由逐格掃描開始，指出從邊界格只能走有效的上下左右鄰格，並理解找到陸地後應立即標記。
- 初稿曾以 `[][]byte` 當 map key、將 byte 與數值 `1` 比較、越界條件反向，以及在 DFS 前先標記起點；經局部提示後自行改用 `map[[2]int]bool` 與正確的標記及四方向遞迴。
- 使用者口述：發現未處理陸地就擴展並標記整座島，因此每島只加一次；時間 O(m*n)，額外空間 O(m*n)。
