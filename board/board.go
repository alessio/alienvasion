package board

import (
	"fmt"
	"math/rand"
)

type board struct{}

// Board defines the
type Board interface {
	// Functional to board creation/initialization
	AddLocation(name string)
	LinkLocations(from, to string, dir Direction) error
	DeployPiece(p Piece) Location
	// Functional to death match simulation
	DestroyLocation(name string)
	HasLinks() bool
	Location(name string) Location
	Locations() []Location
	MovePiece(p Piece) error
	Pieces() []Piece
	WhereIs(p Piece) Location
}

type worldMap struct {
	locations map[string]Location
	pieces    map[Piece]Location
}

func NewBoard() *worldMap {
	return &worldMap{
		locations: make(map[string]Location),
		pieces:    make(map[Piece]Location),
	}
}

func (w *worldMap) AddLocation(name string) { w.locations[name] = NewLocation(name) }
func (w *worldMap) DeployPiece(p Piece) Location {
	locations := w.Locations()
	if len(locations) == 0 {
		panic("no locations left")
	}
	l := w.Locations()[rand.Intn(len(locations))]
	l.AddPiece(p)
	w.pieces[p] = l
	return l
}

func (w *worldMap) DestroyLocation(name string) {
	location, ok := w.locations[name]
	if !ok {
		panic(fmt.Sprintf("location %q does not exist", name))
	}
	location.DestroyLinks()
	pieces := location.Pieces()
	for _, p := range pieces {
		delete(w.pieces, p)
	}
	delete(w.locations, name)
}

func (w *worldMap) LinkLocations(from, to string, dir Direction) error {
	fromLocation, ok := w.locations[from]
	if !ok {
		return fmt.Errorf("location %q does not exist", from)
	}
	toLocation, ok := w.locations[to]
	if !ok {
		return fmt.Errorf("location %q does not exist", to)
	}
	fromLocation.LinkToNeighbour(dir, toLocation)
	toLocation.LinkToNeighbour(OppositeDirection(dir), fromLocation)
	return nil
}

func (w *worldMap) HasLinks() bool {
	for _, location := range w.locations {
		if len(location.ReachableNeighbours()) > 0 {
			return true
		}
	}
	return false
}

func (w *worldMap) Pieces() []Piece {
	pieces := []Piece{}
	for p := range w.pieces {
		pieces = append(pieces, p)
	}
	return pieces
}

func (w *worldMap) Location(name string) Location {
	if location, ok := w.locations[name]; ok {
		return location
	}
	return nil
}

func (w *worldMap) Locations() []Location {
	locations := []Location{}
	for _, location := range w.locations {
		locations = append(locations, location)
	}
	return locations
}

func (w *worldMap) MovePiece(p Piece) error {
	pieceLocation, ok := w.pieces[p]
	if !ok {
		panic(fmt.Sprintf("piece %q does not exist", string(p)))
	}
	next := pickRandomNeighbour(pieceLocation)
	if next == nil {
		return fmt.Errorf("%s trapped at %s", string(p), pieceLocation.Name())
	}
	pieceLocation.RemovePiece(p)
	next.AddPiece(p)
	w.pieces[p] = next
	return nil
}

func (w *worldMap) WhereIs(p Piece) Location {
	return w.pieces[p]
}

func pickRandomNeighbour(location Location) Location {
	neighbours := location.ReachableNeighbours()
	count := len(neighbours)
	if count == 0 {
		return nil
	}
	return neighbours[rand.Intn(count)]
}
