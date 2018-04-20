package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_worldMap_LinkLocations(t *testing.T) {
	board1 := NewBoard()
	board1.AddLocation("a")
	board1.AddLocation("b")
	board1.AddLocation("c")
	board1.AddLocation("d")
	board1.LinkLocations("a", "b", South)
	board1.LinkLocations("b", "c", East)
	board1.LinkLocations("b", "d", West)
	assert.Equal(t, board1.Location("a").Neighbour(South).Name(), "b")
	require.NotNil(t, board1.Location("a"))
	require.NotNil(t, board1.Location("b"))
	require.NotNil(t, board1.Location("c"))
	require.NotNil(t, board1.Location("d"))
	assert.Equal(t, board1.Location("a").Neighbour(South).Name(), board1.Location("b").Name())
	assert.Equal(t, board1.Location("a").Neighbour(North), nil)
	assert.Equal(t, board1.Location("a").Neighbour(West), nil)
	assert.Equal(t, board1.Location("a").Neighbour(East), nil)
	assert.Equal(t, board1.Location("b").Neighbour(North), board1.Location("a"))
	assert.Equal(t, board1.Location("b").Neighbour(South), nil)
	assert.Equal(t, board1.Location("b").Neighbour(West), board1.Location("d"))
	assert.Equal(t, board1.Location("b").Neighbour(East), board1.Location("c"))

	assert.Equal(t, board1.Location("c").Neighbour(West), board1.Location("b"))
	assert.Equal(t, board1.Location("c").Neighbour(North), nil)
	assert.Equal(t, board1.Location("c").Neighbour(South), nil)
	assert.Equal(t, board1.Location("c").Neighbour(East), nil)

	assert.Equal(t, board1.Location("d").Neighbour(East), board1.Location("b"))
	assert.Equal(t, board1.Location("d").Neighbour(North), nil)
	assert.Equal(t, board1.Location("d").Neighbour(South), nil)
	assert.Equal(t, board1.Location("d").Neighbour(West), nil)
}
