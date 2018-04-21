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
	filename string
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Mad aliens are about to invade the earth and this
program aims to simulate the invasion.

Other than the flags described below, this program reads a map from
STDIN. A map contains the names of cities in the non-existent world of
X, one city each line. The city name comes first, followed by 1-4
directions (north, south, east, or west). Each one represents a
road to another city that lies in that direction.
The city and each of the pairs are separated by a single space, and
the directions are separated from their respective cities with an
equals (=) sign.

N aliens are created and given random name, where N is specified as
an optional command-line argument. These aliens start out at random
places on the map, and wander around randomly, following links. Each
iteration, the aliens can travel in any of the directions leading out
of a city.

When multiple aliens end up in the same place, they fight, and in the
process kill each other and destroy the city. When a city is destroyed,
it is removed from the map, and so are any roads that lead into or out
of it. Once a city is destroyed, aliens can no longer travel to or
through it. This may lead to aliens getting "trapped".

`)
		flag.PrintDefaults()
	}
	flag.IntVar(&naliens, "naliens", DefaultNAliens, "generate N aliens")
	flag.IntVar(&maxmoves, "maxmoves", DefaultMaxMoves, "max number of moves each alien can make")
	flag.StringVar(&filename, "file", "", "read map from file instead of STDIN")
	log.SetPrefix(fmt.Sprintf("%s: ", path.Base(os.Args[0])))
	//	log.SetFlags(0)
}

func main() {
	var fp io.Reader = os.Stdin
	flag.Parse()
	validateFlags()
	// Initialise the board
	if filename != "" {
		fp = mustOpenFile(filename)
	}
	b := mustParseBoard(fp)
	print(board.EncodeBoard(b))
	//Initialise the simulation
	dm := deathmatch.NewDeathMatch(b, maxmoves)
	dm.KickOff(naliens)
	for dm.ExecuteTurn() {
		//
	}
	fmt.Println("What is left of the world?")
	print(board.EncodeBoard(b))
}

func mustOpenFile(filename string) io.Reader {
	fp, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return fp
}

func mustParseBoard(reader io.Reader) *board.Board {
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
