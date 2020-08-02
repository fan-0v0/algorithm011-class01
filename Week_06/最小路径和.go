func minPathSum(grid [][]int) int {
	np := make([][]int, len(grid))
	for i := range np {
		np[i] = make([]int, len(grid[0]))
	}
	np[0][0] = grid[0][0]
	for i := 1; i < len(grid); i++ {
		np[i][0] = np[i-1][0] + grid[i][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		np[0][j] = np[0][j-1] + grid[0][j]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			np[i][j] = helper(np[i-1][j], np[i][j-1]) + grid[i][j]
		}
	}
	return np[len(grid)-1][len(grid[0])-1]
}

func helper(a, b int) int {
	if a > b {
		a = b
	}
	return a
}