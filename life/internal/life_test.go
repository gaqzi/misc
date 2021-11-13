package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	gol "not-so-random/life/internal"
)

func TestNeighbours(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]int
		row      int
		col      int
		expected int
	}{
		{
			name:     "no neighbours is 0",
			input:    [][]int{{0}},
			row:      0,
			col:      0,
			expected: 0,
		},
		{
			name:     "1 neighbour to the right",
			input:    [][]int{{0, 1}},
			row:      0,
			col:      0,
			expected: 1,
		},
		{
			name:     "1 neighbour to the left",
			input:    [][]int{{1, 0}},
			row:      0,
			col:      1,
			expected: 1,
		},
		{
			name: "1 neighbour above",
			input: [][]int{
				{1},
				{0},
			},
			row:      1,
			col:      0,
			expected: 1,
		},
		{
			name: "1 neighbour above: left",
			input: [][]int{
				{1, 0},
				{0, 0},
			},
			row:      1,
			col:      1,
			expected: 1,
		},
		{
			name: "1 neighbour above: right",
			input: [][]int{
				{0, 1},
				{0, 0},
			},
			row:      1,
			col:      0,
			expected: 1,
		},
		{
			name: "1 neighbour below",
			input: [][]int{
				{0},
				{1},
			},
			row:      0,
			col:      0,
			expected: 1,
		},
		{
			name: "1 neighbour below left",
			input: [][]int{
				{0, 0},
				{1, 0},
			},
			row:      0,
			col:      1,
			expected: 1,
		},
		{
			name: "1 neighbour below right",
			input: [][]int{
				{0, 0},
				{0, 1},
			},
			row:      0,
			col:      0,
			expected: 1,
		},
		{
			name: "3x3 full grid",
			input: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			row:      1,
			col:      1,
			expected: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, gol.Neighbours(tc.input, tc.row, tc.col))
		})
	}
}

func TestIsAlive(t *testing.T) {
	testCases := []struct {
		name       string
		alive      int
		neighbours int
		expected   int
	}{
		{
			name:       "2: is alive and stays alive",
			alive:      1,
			neighbours: 2,
			expected:   1,
		},
		{
			name:       "2: is dead and stays dead",
			alive:      1,
			neighbours: 2,
			expected:   1,
		},
		{
			name:       "3: is alive and stays alive",
			alive:      1,
			neighbours: 3,
			expected:   1,
		},
		{
			name:       "3: is dead and is given life",
			alive:      0,
			neighbours: 3,
			expected:   1,
		},
		{
			name:       "1: alive and becomes dead",
			alive:      1,
			neighbours: 1,
			expected:   0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, gol.IsAlive(tc.alive, tc.neighbours))
		})
	}

	t.Run("values above 4 dies", func(t *testing.T) {
		maxNeighbours := 9 // in a grid you can only have 8 neighbours with the center point being 0
		for i := 4; i <= maxNeighbours; i++ {
			assert.Equal(t, 0, gol.IsAlive(1, i), "expected %d to be dead", i)
		}
	})
}
