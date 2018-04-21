package deathmatch

import (
	"testing"

	"github.com/alessio/alienvasion/board"
	"github.com/stretchr/testify/assert"
)

func TestDeathMatch_isGameOver(t *testing.T) {
	var d *DeathMatch
	makeBoard := func() *board.Board {
		b := board.NewBoard()
		b.AddLocation("a")
		b.AddLocation("b")
		b.AddLocation("c")
		b.AddLocation("d")
		b.LinkLocations(board.Location("a"), board.Location("b"), board.Direction(board.East))
		b.LinkLocations(board.Location("b"), board.Location("c"), board.Direction(board.South))
		b.LinkLocations(board.Location("c"), board.Location("d"), board.Direction(board.West))
		b.LinkLocations(board.Location("d"), board.Location("a"), board.Direction(board.North))
		return b
	}
	// No locations on the board
	d = NewDeathMatch(board.NewBoard(), 10)
	assert.Panics(t, func() { d.KickOff(10) })

	// No pieces on the board
	d = NewDeathMatch(makeBoard(), 10)
	d.KickOff(0)
	assert.NotNil(t, d.isGameOver())
	// Add one alien
	d.KickOff(1)
	// Still not enough aliens
	assert.NotNil(t, d.isGameOver())

	// All locations destroyed
	b := makeBoard()
	d = NewDeathMatch(b, 10)
	for _, n := range []string{"a", "b", "c", "d"} {
		b.Destroy(board.Location(n), []board.Piece{})
	}
	assert.NotNil(t,
		d.isGameOver())
}
