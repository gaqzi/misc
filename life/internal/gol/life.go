package gol

// Neighbours returns how many adjacent cells are alive
func Neighbours(grid [][]int, row, col int) int {
	var total int

	if right := col + 1; right < len(grid[row]) {
		total += grid[row][right]
	}

	if left := col - 1; left >= 0 {
		total += grid[row][left]
	}

	if above := row - 1; above >= 0 {
		if left := col - 1; left >= 0 {
			total += grid[above][left]
		}

		total += grid[above][col]

		if right := col + 1; right < len(grid[above]) {
			total += grid[above][right]
		}
	}

	if below := row + 1; below < len(grid) {
		if left := col - 1; left >= 0 {
			total += grid[below][left]
		}

		total += grid[below][col]

		if right := col + 1; right < len(grid[below]) {
			total += grid[below][right]
		}
	}

	return total
}

func IsAlive(neighbours int) int {
	switch neighbours {
	case 2, 3:
		return 1
	default:
		return 0
	}
}
