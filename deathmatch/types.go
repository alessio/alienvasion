package deathmatch

import (
	"github.com/alessio/alienvasion/board"
)

// DeathMatch defines the interface of a death match.
type DeathMatch interface {
	Board() board.Board
	KickOff(npieces int)
	ExecuteTurn() error
}
