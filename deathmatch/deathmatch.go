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
	log.Print("Kick off!")
	for i := 0; i < npieces; i++ {
		name := fmt.Sprintf("%s-%d", generateSillyName(), i)
		location := d.board.DeployPiece(board.Piece(name))
		log.Printf(" - %q deployed on %q", name, location.Name())
	}
}

func (d *deathmatch) ExecuteTurn() bool {
	defer func() { d.turn++ }()
	log.Printf("]-[ Turn #%d ]-[", d.turn)
	if err := d.isGameOver(); err != nil {
		log.Println(err.Error())
		return false
	}
	d.movePieces()
	d.fight()
	return true
}

func (d *deathmatch) isGameOver() error {
	if d.turn >= d.maxmoves {
		return fmt.Errorf("reached maximum number of moves")
	}
	if len(d.board.Locations()) == 0 {
		return errors.New("every man for himself! the world is collapsing")
	}
	if !d.board.HasLinks() {
		return errors.New("game over - all aliens are trapped")
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
			log.Print(err.Error())
		}
	}
}

func (d *deathmatch) fight() {
	for _, location := range d.board.Locations() {
		if len(location.Pieces()) >= 2 {
			d.board.DestroyLocation(location.Name())
			log.Printf("%q has been destroyed by %s", location.Name(), location.Pieces())
		}
	}
}

func generateSillyName() string {
	return petname.Generate(3, "-")
}
