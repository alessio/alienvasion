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
		pieceName := fmt.Sprintf("%s-%d", generateSillyName(), i)
		d.board.DeployRandPiece(pieceName)
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
	if err := d.movePieces(); err != nil {
		return err
	}
	return d.fight()
}

func (d *deathmatch) isGameOver() error {
	if len(d.board.Locations()) == 0 {
		return errors.New("every man for himself! the world is collapsing")
	}
	if !d.board.HasLinks() {
		return errors.New("game over: all aliens are stuck")
	}
	return nil
}

func (d *deathmatch) movePieces() error {
	// First, move aliens
	for _, piece := range d.board.Pieces() {
		if err := piece.Wander(); err != nil {
			log.Printf("%s is stuck at %s", piece.Name(), piece.Location().Name())
		}
	}
	return nil
}

func (d *deathmatch) fight() error {
	for _, location := range d.board.Locations() {
		if len(location.Pieces()) >= 2 {
			d.board.DestroyLocation(location)
		}
	}
	if len(d.board.Pieces()) < 2 {
		return errors.New("shut down the internet, folks! We have a winner")
	}
	return nil
}

func generateSillyName() string {
	return petname.Generate(3, "-")
}
