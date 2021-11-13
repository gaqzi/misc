package life_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"misc/life"
)

/*

1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
2. Any live cell with more than three live neighbours dies, as if by overcrowding.
3. Any live cell with two or three live neighbours lives on to the next generation.
4. Any dead cell with exactly three live neighbours becomes a live cell.
 */
func Test_Integration(t *testing.T) {
	input := `4 8
........
....*...
...**...
.....*..
`
	expected := `4 8
........
...**...
...***..
....*...
`

	game, err := life.Game(input)
	require.NoError(t, err)

	require.Equal(t, expected, life.Evolve(game).String())
}

func Test_Game(t *testing.T) {
	// These tests do kind of a lot in hindsight, but it also feels right,
	// because parsing and outputting the canonical string format is important.

	testCases := []struct {
		name     string
		input    string
		expected life.World
	}{
		{
			name:     "empty World",
			input:    "0 0\n",
			expected: life.World{},
		},
		{
			name:     "one grid World",
			input:    "1 1\n.\n",
			expected: life.World{{0}},
		},
		{
			name:  "2x2 grid World with alive and dead cells",
			input: "2 2\n.*\n**\n",
			expected: life.World{
				{0, 1},
				{1, 1},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			game, err := life.Game(tc.input)
			require.NoError(t, err)

			assert.Equal(t, tc.expected, game, "expected the game world to be parsed correctly")
			assert.Equal(t, tc.input, game.String(), "expected the world to output the same canonical string")
		})
	}
}

func Test_Evolve(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Cell with only one neighbour dies",
			input:    "1 2\n.*\n",
			expected: "1 2\n..\n",
		},
		{
			name:     "Cell with two neighbour stays alive",
			input:    "1 3\n***\n",
			expected: "1 3\n.*.\n",
		},
		{
			name:     "Cell with three neighbour comes alive",
			input:    "2 3\n*.*\n.*.\n",
			expected: "2 3\n.*.\n.*.\n",
		},
		{
			name:     "Cell with three neighbour stays alive",
			input:    "2 3\n***\n.*.\n",
			expected: "2 3\n***\n***\n",
		},
		{
			name:     "Cell with four neighbour dies",
			input:    "2 3\n***\n***\n",
			expected: "2 3\n*.*\n*.*\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			game, err := life.Game(tc.input)
			require.NoError(t, err)

			require.Equal(t, tc.expected, life.Evolve(game).String())
		})
	}
}
