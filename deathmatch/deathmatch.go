package deathmatch

import (
	"errors"
	"fmt"
	"log"

	"github.com/alessio/alienvasion/board"
	petname "github.com/dustinkirkland/golang-petname"
)

func NewDeathMatch(b board.Board, maxmoves int) DeathMatch {
	return &deathmatch{board: b, maxmoves: maxmoves, turn: 0}
}

type deathmatch struct {
	board    board.Board
	maxmoves int
	turn     int
}

func (d *deathmatch) Board() board.Board {
	return d.board
}

func (d *deathmatch) KickOff(npieces int) {
	for i := 0; i < npieces; i++ {
		d.board.DeployPiece(board.Piece(
			fmt.Sprintf("%s-%d", generateSillyName(), i)))
	}
}

func (d *deathmatch) ExecuteTurn() error {
	defer func() { d.turn++ }()
	if d.turn >= d.maxmoves {
		return fmt.Errorf("reached maximum number of moves")
	}
	if err := d.isGameOver(); err != nil {
		return err
	}
	d.movePieces()
	d.fight()
	return nil
}

func (d *deathmatch) isGameOver() error {
	if len(d.board.Locations()) == 0 {
		return errors.New("every man for himself! the world is collapsing")
	}
	if !d.board.HasLinks() {
		return errors.New("game over - all aliens are stuck")
	}
	if len(d.board.Pieces()) < 2 {
		return errors.New("war is over")
	}
	return nil
}

func (d *deathmatch) movePieces() {
	// First, move aliens
	for _, piece := range d.board.Pieces() {
		if err := d.board.MovePiece(piece); err != nil {
			log.Printf("piece.Wander: %s", err.Error())
		}
	}
}

func (d *deathmatch) fight() {
	for _, location := range d.board.Locations() {
		if len(location.Pieces()) >= 2 {
			d.board.DestroyLocation(location.Name())
		}
	}
}

func generateSillyName() string {
	return petname.Generate(3, "-")
}
