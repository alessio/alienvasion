package board

import (
	"fmt"
	"math/rand"
)

const (
	North = Direction("north")
	South = Direction("south")
	West  = Direction("west")
	East  = Direction("east")
)

type Direction string

func (d Direction) Validate() bool { return d == North || d == South || d == West || d == East }
func (d Direction) Opposite() Direction {
	switch d {
	case North:
		return South
	case South:
		return North
	case West:
		return East
	case East:
		return West
	}
	panic("invalid direction")
}

type Location string
type Piece string

type Board struct {
	links  map[Location]map[Direction]Location
	pieces map[Piece]Location
}

func New() *Board {
	return &Board{
		links:  make(map[Location]map[Direction]Location),
		pieces: make(map[Piece]Location),
	}
}

func (b *Board) Pieces() map[Piece]Location { return b.pieces }
func (b *Board) AddLocation(l Location) {
	if _, ok := b.links[l]; !ok {
		b.links[l] = make(map[Direction]Location)
	}
}

func (b *Board) LinkLocations(l1, l2 Location, dir Direction) error {
	if l1 == l2 {
		return fmt.Errorf("loop detected")
	}
	oppositeDir := dir.Opposite()
	b.AddLocation(l2)
	if loc, ok := b.links[l1][dir]; ok && loc != l2 {
		return fmt.Errorf("couldn't overwrite link between %s and %s (%s)", l1, loc, dir)
	}
	if loc, ok := b.links[l2][oppositeDir]; ok && loc != l1 {
		return fmt.Errorf("couldn't overwrite link between %s and %s (%s)", l2, loc, oppositeDir)
	}
	b.links[l1][dir] = l2
	b.links[l2][oppositeDir] = l1
	return nil
}

func (b *Board) Deploy(l Location, p Piece) {
	if _, ok := b.pieces[p]; ok {
		panic("%s already exists")
	}
	b.pieces[p] = l
}

func (b *Board) PieceNeighbourLocations(p Piece) []Location {
	loc, ok := b.pieces[p]
	if !ok {
		panic("piece slipped off the board")
	}
	dirmap := b.links[loc]
	if dirmap == nil {
		return nil
	}
	locations := []Location{}
	for _, loc := range dirmap {
		locations = append(locations, loc)
	}
	return locations
}

func (b *Board) Wander(p Piece) (Location, error) {
	locations := b.PieceNeighbourLocations(p)
	if locations == nil || len(locations) == 0 {
		return "", fmt.Errorf("%s is trapped at %s", p, b.pieces[p])
	}
	b.pieces[p] = locations[rand.Intn(len(locations))]
	return b.pieces[p], nil
}
func (b *Board) move(p Piece, l Location) {
	currentLoc := b.pieces[p]
	if !b.AreNeighbours(currentLoc, l) {
		panic(fmt.Sprintf("%s attempted to make an invalid move from %s to %s", p, currentLoc, l))
	}
	b.pieces[p] = l
}

func (b *Board) Neighbours(l Location) []Location {
	dirmap, ok := b.links[l]
	if !ok {
		return nil
	}
	neighbours := []Location{}
	for _, adj := range dirmap {
		neighbours = append(neighbours, adj)
	}
	return neighbours
}

func (b *Board) AreNeighbours(l1, l2 Location) bool {
	dirmap, ok := b.links[l1]
	if !ok {
		return false
	}
	for _, neigh := range dirmap {
		if neigh == l2 {
			return true
		}
	}
	return false
}

func (b *Board) Locations() (locations []Location) {
	for loc := range b.links {
		locations = append(locations, loc)
	}
	return
}

func (b *Board) PiecesByLocation() map[Location][]Piece {
	piecesByLocation := make(map[Location][]Piece)
	for p, loc := range b.pieces {
		if _, ok := piecesByLocation[loc]; !ok {
			piecesByLocation[loc] = []Piece{}
		}
		piecesByLocation[loc] = append(piecesByLocation[loc], p)
	}
	return piecesByLocation
}

func (b *Board) Destroy(l Location, pieces []Piece) {
	for _, p := range pieces {
		delete(b.pieces, p)
	}
	for dir, neighbour := range b.links[l] {
		delete(b.links[neighbour], dir.Opposite())
	}
	delete(b.links, l)
}
