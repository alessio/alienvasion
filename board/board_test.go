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

func Test_worldMap_LinkLocations_Simmetry(t *testing.T) {
	board1 := NewBoard()
	board1.AddLocation("a")
	board1.AddLocation("b")
	board1.LinkLocations("a", "b", West)
	board2 := NewBoard()
	board2.AddLocation("a")
	board2.AddLocation("b")
	board2.LinkLocations("b", "a", East)
	assert.Equal(
		t,
		board1.Location("a").Neighbour(West).Name(),
		board2.Location("a").Neighbour(West).Name(),
	)
	assert.Equal(
		t,
		board1.Location("b").Neighbour(East).Name(),
		board2.Location("b").Neighbour(East).Name(),
	)
}

func Test_worldMap_DestroyLocation(t *testing.T) {
	board1 := NewBoard()
	board1.AddLocation("a")
	board1.AddLocation("b")
	board1.AddLocation("c")
	board1.DeployPiece("x")
	board1.DeployPiece("y")
	board1.DeployPiece("w")
	board1.DeployPiece("z")
	// a <==> b <==> c
	board1.LinkLocations("a", "b", West)
	board1.LinkLocations("b", "c", West)
	require.Equal(t, board1.Location("a").Neighbour(West).Name(), "b")
	require.Equal(t, board1.Location("b").Neighbour(West).Name(), "c")
	require.Equal(t, board1.Location("c").Neighbour(East).Name(), "b")
	require.Equal(t, board1.Location("b").Neighbour(East).Name(), "a")
	require.Equal(t, len(board1.Pieces()), 4)

	nDeadPieces := len(board1.Location("b").Pieces())
	board1.DestroyLocation("b")
	assert.Nil(t, board1.Location("a").Neighbour(West))
	assert.Nil(t, board1.Location("c").Neighbour(East))
	assert.Equal(t, len(board1.Pieces()), 4-nDeadPieces)
}

func Test_worldMap_DeployPiece(t *testing.T) {
	names := []string{"a", "b", "c"}
	board1 := NewBoard()
	// Panic expected, board is empty
	assert.Panics(t, func() { board1.DeployPiece("z") })
	// Add locations
	for _, n := range names {
		board1.AddLocation(n)
	}
	board1.DeployPiece("x")
	assert.Equal(t, board1.Pieces(), []Piece{Piece("x")})
}

func Test_worldMap_MovePiece(t *testing.T) {
	board1 := NewBoard()
	board1.AddLocation("a")
	board1.DeployPiece("x")
	require.Equal(t, board1.Pieces(), []Piece{Piece("x")})
	require.Equal(t, board1.Location("a").Pieces(), []Piece{Piece("x")})
	// Add a location
	board1.AddLocation("b")
	board1.LinkLocations("a", "b", West)
	assert.Equal(t, board1.Pieces(), []Piece{Piece("x")})
	assert.Equal(t, board1.Location("a").Pieces(), []Piece{Piece("x")})
	assert.Equal(t, board1.Location("b").Pieces(), []Piece{})
	// assert.Equal(t, board1.Pieces(), []Piece{Piece("x")})
	board1.MovePiece("x")
	assert.Equal(t, board1.Pieces(), []Piece{Piece("x")})
	assert.Equal(t, board1.Location("a").Pieces(), []Piece{})
}
