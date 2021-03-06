package internal

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

// IsAlive returns whether this value should be alive or dead based on the number of neighbours it has.
// Following the Game of Life rules.
func IsAlive(alive, neighbours int) int {
	switch {
	case neighbours == 2 && alive == 1:
		return 1
	case neighbours == 3:
		return 1
	default:
		return 0
	}
}
