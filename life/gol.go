package life

import (
	"fmt"
	"strconv"
	"strings"

	"not-so-random/life/internal/gol"
)

type World struct {
	Grid [][]int
}

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
	game := make([][]int, 0, rows)

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

	return World{Grid: game}, nil
}

func (w World) String() string {
	if len(w.Grid) == 0 {
		return ""
	}

	var output string
	output += fmt.Sprintf("%d %d\n", len(w.Grid), len(w.Grid[0]))

	for _, row := range w.Grid {
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

func Evolve(w World) World {
	grid := make([][]int, 0, len(w.Grid))

	for row := 0; row < len(w.Grid); row++ {
		cells := make([]int, 0, len(w.Grid[row]))

		for col := 0;col < len(w.Grid[row]);col++ {
			cells = append(cells, gol.IsAlive(w.Grid[row][col], gol.Neighbours(w.Grid, row, col)))
		}

		grid = append(grid, cells)
	}

	return World{Grid: grid}
}
