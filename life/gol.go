package life

import (
	"fmt"
	"strconv"
	"strings"
)

type World struct {
	grid [][]int
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

	return World{grid: game}, nil
}

func (w World) String() string {
	if len(w.grid) == 0 {
		return ""
	}

	var output string
	output += fmt.Sprintf("%d %d\n", len(w.grid), len(w.grid[0]))

	for _, row := range w.grid {
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
	return World{grid: [][]int{{0, 0}}}
}
