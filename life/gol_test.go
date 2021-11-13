package life_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"not-so-random/life"
)

/*

1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
2. Any live cell with more than three live neighbours dies, as if by overcrowding.
3. Any live cell with two or three live neighbours lives on to the next generation.
4. Any dead cell with exactly three live neighbours becomes a live cell.
 */
func Test_Integration(t *testing.T) {
	t.Skip("wait until pieces built up")
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
	t.Run("empty World", func(t *testing.T) {
		game, err := life.Game("0 0")
		require.NoError(t, err)

		require.Equal(t, "", game.String())
	})

	t.Run("one grid World", func(t *testing.T) {
		game, err := life.Game("1 1\n.")
		require.NoError(t, err)

		require.Equal(t, "1 1\n.\n", game.String())
	})

	t.Run("2x2 grid World with alive and dead cells", func(t *testing.T) {
		game, err := life.Game("2 2\n.*\n**\n")
		require.NoError(t, err)

		require.Equal(t, "2 2\n.*\n**\n", game.String())
	})
}

func Test_Evolve(t *testing.T) {

}