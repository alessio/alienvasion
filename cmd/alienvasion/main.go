package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"github.com/alessio/alienvasion/board"
	"github.com/alessio/alienvasion/deathmatch"
)

const (
	// DefaultNAliens is the default number of aliens that are
	// generated if no value is supplied by the user.
	DefaultNAliens = 8
	// DefaultMaxMoves is the default maximum number of moves
	// if no value is supplied by the user.
	DefaultMaxMoves = 10000
)

var (
	naliens  int
	maxmoves int
)

func init() {
	flag.IntVar(&naliens, "naliens", DefaultNAliens, "generate N aliens")
	flag.IntVar(&maxmoves, "maxmoves", DefaultMaxMoves, "max number of moves each alien can make")
	log.SetPrefix(fmt.Sprintf("%s: ", path.Base(os.Args[0])))
	log.SetFlags(0)
}

func main() {
	flag.Parse()
	validateFlags()
	// Initialise the board
	b := mustParseBoard(os.Stdin)
	fmt.Println("Board generated and aliens deployed:")
	board.PrintBoard(b)
	validateAliensAgainstBoardSize(len(b.Locations()))
	// Initialise the simulation
	dm := deathmatch.NewDeathMatch(b, maxmoves)
	dm.KickOff(naliens)
	for {
		if err := dm.ExecuteTurn(); err != nil {
			fmt.Printf("%s", err.Error())
			break
		}
	}
}

func mustParseBoard(reader io.Reader) board.Board {
	b, err := board.ParseBoard(reader)
	if err != nil {
		log.Fatalf("board.ParseBoard: %v", err)
	}
	return b
}

func validateFlags() {
	if naliens < 2 {
		log.Fatal("the number of aliens must be greater than 1.")
	}
	if maxmoves < 0 {
		log.Fatal("the maximum number of moves must be equal or greater than 0.")
	}
}

func validateAliensAgainstBoardSize(boardSize int) {
	if naliens > boardSize {
		log.Fatalf(
			"the number of aliens (%d) must be equal or smaller than the number of cities (%d).",
			naliens, boardSize)
	}
}
