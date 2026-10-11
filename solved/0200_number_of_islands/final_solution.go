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
