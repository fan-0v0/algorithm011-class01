func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				DFS(i, j, grid)
				count++
			}
		}
	}
	return count
}

func DFS(i, j int, grid [][]byte) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	DFS(i-1, j, grid)
	DFS(i+1, j, grid)
	DFS(i, j+1, grid)
	DFS(i, j-1, grid)
}