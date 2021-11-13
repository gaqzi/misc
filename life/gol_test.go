package life_test

import (
	"testing"

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

	t.Run("empty World", func(t *testing.T) {
		game, err := life.Game("0 0")
		require.NoError(t, err)

		require.Equal(t, life.World{}, game)
		require.Equal(t, "", game.String())
	})

	t.Run("one grid World", func(t *testing.T) {
		game, err := life.Game("1 1\n.")
		require.NoError(t, err)

		require.Equal(t, life.World{{0}}, game)
		require.Equal(t, "1 1\n.\n", game.String())
	})

	t.Run("2x2 grid World with alive and dead cells", func(t *testing.T) {
		game, err := life.Game("2 2\n.*\n**\n")
		require.NoError(t, err)

		require.Equal(
			t,
			life.World{
				{0, 1},
				{1, 1},
			},
			game,
		)
		require.Equal(t, "2 2\n.*\n**\n", game.String())
	})
}

func Test_Evolve(t *testing.T) {
	t.Run("Cell with only one neighbour dies", func(t *testing.T) {
		game, err := life.Game("1 2\n.*\n")
		require.NoError(t, err)

		require.Equal(t, "1 2\n..\n", life.Evolve(game).String())
	})

	t.Run("Cell with two neighbour stays alive", func(t *testing.T) {
		game, err := life.Game("1 3\n***\n")
		require.NoError(t, err)

		require.Equal(t, "1 3\n.*.\n", life.Evolve(game).String())
	})

	t.Run("Cell with three neighbour comes alive", func(t *testing.T) {
		game, err := life.Game("2 3\n*.*\n.*.\n")
		require.NoError(t, err)

		require.Equal(t, "2 3\n.*.\n.*.\n", life.Evolve(game).String())
	})

	t.Run("Cell with three neighbour stays alive", func(t *testing.T) {
		game, err := life.Game("2 3\n***\n.*.\n")
		require.NoError(t, err)

		require.Equal(t, "2 3\n***\n***\n", life.Evolve(game).String())
	})

	t.Run("Cell with four neighbour dies", func(t *testing.T) {
		game, err := life.Game("2 3\n***\n***\n")
		require.NoError(t, err)

		require.Equal(t, "2 3\n*.*\n*.*\n", life.Evolve(game).String())
	})
}
