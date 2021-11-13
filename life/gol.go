package life

import (
	"fmt"
	"strconv"
	"strings"

	gol "misc/life/internal"
)

// World represents a Game of Life world and can output itself in the canonical string format.
type World [][]int

// Game parses a canonical string into a World.
func Game(input string) (World, error) {
	lines := strings.Split(input, "\n")

	var size []int
	for _, item := range strings.SplitN(lines[0], " ", 2) {
		n, err := strconv.Atoi(item)
		if err != nil {
			return World{}, fmt.Errorf("failed to parse size '%s': %w", lines[0], err)
		}

		size = append(size, n)
	}

	rows, cols := size[0], size[1]
	game := make(World, 0, rows)

	for i := 0; i < rows; i++ {
		row := strings.Split(lines[i+1], "")
		cells := make([]int, cols)

		for j, v := range row {
			switch v {
			case ".":
				cells[j] = 0
			case "*":
				cells[j] = 1
			default:
				return World{}, fmt.Errorf("unknown value on input cell %dx%d (%s): %s", i, j, lines[i+1], v)
			}
		}

		game = append(game, cells)
	}

	return game, nil
}

// String outputs the World as a canonical string.
func (w World) String() string {
	if len(w) == 0 {
		return "0 0\n"
	}

	var output string
	output += fmt.Sprintf("%d %d\n", len(w), len(w[0]))

	for _, row := range w {
		for _, v := range row {
			switch v {
			case 0:
				output += "."
			case 1:
				output += "*"
			default:
				output += "!"
			}
		}
		output += "\n"
	}

	return output
}

// Evolve runs the Game of Life on each cell of the world creating the next generation.
func Evolve(w World) World {
	grid := make(World, 0, len(w))

	for row := 0; row < len(w); row++ {
		cells := make([]int, 0, len(w[row]))

		for col := 0;col < len(w[row]);col++ {
			cells = append(cells, gol.IsAlive(w[row][col], gol.Neighbours(w, row, col)))
		}

		grid = append(grid, cells)
	}

	return grid
}
