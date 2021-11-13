package gol

import "not-so-random/life"

// Neighbours returns how many adjacent cells are alive
func Neighbours(w life.World, row, col int) int {
	var total int

	if right := col + 1; right < len(w.Grid[row]) {
		total += w.Grid[row][right]
	}

	if left := col - 1; left >= 0 {
		total += w.Grid[row][left]
	}

	if above := row - 1; above >= 0 {
		total += w.Grid[above][col]
	}

	if below := row + 1; below < len(w.Grid) {
		total += w.Grid[below][col]
	}

	return total
}
