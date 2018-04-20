package board

const (
	North = "north"
	South = "south"
	West  = "west"
	East  = "east"
)

type Piece string
type Direction string

var oppositeDirectionsMap = map[Direction]Direction{
	North: South,
	South: North,
	West:  East,
	East:  West,
}

type Location interface {
	Name() string
	Pieces() []Piece
	ReachableNeighbours() []Location
	Neighbour(dir Direction) Location
	LinkToNeighbour(dir Direction, target Location)
	DestroyLinks()
	AddPiece(p Piece)
	RemovePiece(p Piece)
}

func (p Piece) Name() string { return string(p) }

type location struct {
	name       string
	pieces     map[Piece]bool
	neighbours map[Direction]Location
}

var _ Location = (*location)(nil)

func NewLocation(name string) Location {
	return &location{
		name:   name,
		pieces: make(map[Piece]bool),
		neighbours: map[Direction]Location{
			North: nil,
			South: nil,
			West:  nil,
			East:  nil,
		},
	}
}

func (l *location) Name() string { return l.name }
func (l *location) Neighbour(dir Direction) Location {
	if n, ok := l.neighbours[dir]; ok {
		return n
	}
	return nil
}
func (l *location) LinkToNeighbour(dir Direction, target Location) {
	t := target.(*location)
	l.neighbours[dir] = t
}
func (l *location) AddPiece(p Piece)    { l.pieces[p] = true }
func (l *location) RemovePiece(p Piece) { delete(l.pieces, p) }
func (l *location) Pieces() []Piece {
	pieces := []Piece{}
	for p := range l.pieces {
		pieces = append(pieces, p)
	}
	return pieces
}
func (l *location) ReachableNeighbours() []Location {
	reachable := []Location{}
	for _, val := range l.neighbours {
		if val != nil {
			reachable = append(reachable, val)
		}
	}
	return reachable
}
func (l *location) DestroyLinks() {
	for dir := range l.neighbours {
		delete(l.neighbours, dir)
	}
}
func (l *location) String() string {
	return l.name
}

// Direction implementation

func (d Direction) String() string              { return string(d) }
func OppositeDirection(dir Direction) Direction { return oppositeDirectionsMap[dir] }
func (dir Direction) Validate() bool {
	_, ok := oppositeDirectionsMap[dir]
	return ok
}
