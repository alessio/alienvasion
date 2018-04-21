package deathmatch

import (
	"errors"
	"fmt"
	"log"
	"math/rand"

	"github.com/alessio/alienvasion/board"

	petname "github.com/dustinkirkland/golang-petname"
)

func NewDeathMatch(b *board.Board, maxmoves int) *DeathMatch {
	return &DeathMatch{board: b, maxmoves: maxmoves, turn: 0}
}

type DeathMatch struct {
	board    *board.Board
	maxmoves int
	turn     int
}

func (d *DeathMatch) Board() *board.Board {
	return d.board
}

func (d *DeathMatch) KickOff(npieces int) {
	log.Print("Kick off!")
	locations := d.board.Locations()
	if locations == nil {
		panic("no locations available")
	}
	for i := 0; i < npieces; i++ {
		randomLoc := locations[rand.Intn(len(locations))]
		name := fmt.Sprintf("%s-%d", generateSillyName(), i)
		d.board.Deploy(randomLoc, board.Piece(name))
		log.Printf(" - %q deployed on %q", name, randomLoc)
	}
}

func (d *DeathMatch) ExecuteTurn() bool {
	defer func() { d.turn++ }()
	log.Printf("]-[ Turn #%d ]-[", d.turn)
	if err := d.isGameOver(); err != nil {
		log.Println(err.Error())
		return false
	}
	d.movePieces()
	d.fight()
	// // print out pieces
	log.Println(" - pieces still alive:", alivePieces(d.board))
	log.Println(" - locations still standing:", stillStandingLocations(d.board))
	return true
}

func stillStandingLocations(b *board.Board) string {
	s := ""
	for _, l := range b.Locations() {
		s += fmt.Sprintf("%s ", l)
	}
	return s
}

func alivePieces(b *board.Board) string {
	s := ""
	for p, loc := range b.Pieces() {
		s += fmt.Sprintf("%v:%s ", p, loc)
	}
	return s
}

func (d *DeathMatch) isGameOver() error {
	if d.turn >= d.maxmoves {
		return fmt.Errorf("reached maximum number of moves")
	}
	if d.board.Locations() == nil {
		return errors.New("all locations have been destroyed")
	}
	switch len(d.board.Pieces()) {
	case 0:
		return errors.New("all pieces have been removed")
	case 1:
		return errors.New("we've got a winner")
	}
	return nil
}

func (d *DeathMatch) movePieces() {
	// First, move aliens
	for p, oldloc := range d.board.Pieces() {
		newloc, err := d.board.Wander(p)
		if err != nil {
			log.Print(err.Error())
		} else {
			log.Printf(" - %s moved from %s to %s", p, oldloc, newloc)
		}
	}
}

func (d *DeathMatch) fight() {
	for loc, pieces := range d.board.PiecesByLocation() {
		if len(pieces) >= 2 {
			log.Printf("%s has been destroyed by %s", loc, pieces)
			d.board.Destroy(loc, pieces)
		}
	}
}

func generateSillyName() string {
	return petname.Generate(3, "-")
}
