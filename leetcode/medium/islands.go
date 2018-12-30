package lcmedium

func numIslands(grid [][]byte) int {
	count := 0
	maxSize := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if string(grid[i][j]) == "1" {
				count++
				currentSize := getConnectedLandCount(grid, i, j)
				if currentSize > maxSize {
					maxSize = currentSize
				}
			}
		}
	}

	return count
}
func getConnectedLandCount(grid [][]byte, i int, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || string(grid[i][j]) != "1" {
		return 0
	}

	grid[i][j] = 48
	size := 1
	/*for row := i - 1; row <= i+1; row++ {
			for col := j - 1; col <= j+1; col++ {
				if col != j || row != i {
					size += getConnectedLandCount(grid, row, col)
				}
			}
	  }*/
	size += getConnectedLandCount(grid, i+1, j)
	size += getConnectedLandCount(grid, i-1, j)
	size += getConnectedLandCount(grid, i, j-1)
	size += getConnectedLandCount(grid, i, j+1)

	return size
}
