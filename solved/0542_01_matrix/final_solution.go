package main

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
