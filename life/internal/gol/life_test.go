package gol_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"not-so-random/life"
	"not-so-random/life/internal/gol"
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			game := life.World{Grid: tc.input}

			assert.Equal(t, tc.expected, gol.Neighbours(game, tc.row, tc.col))
		})
	}
}

func TestIsAlive(t *testing.T) {
	testCases := []struct {
		name       string
		neighbours int
		expected   bool
	}{
		{
			name: "2: alive",
			neighbours: 2,
			expected: true,
		},
		{
			name: "3: alive",
			neighbours: 3,
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, gol.IsAlive(tc.neighbours))
		})
	}
}
