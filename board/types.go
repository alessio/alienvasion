package board

type Board interface {
	// Functional to board creation
	AddLocation(name string)
	LinkLocations(from, to string, dir Direction) error
	// Functional to death match simulation
	DestroyLocation(location Location)
	DeployRandPiece(name string) Piece
	HasLinks() bool
	Locations() []Location
	Pieces() []Piece
}

type Piece interface {
	Name() string
	Location() Location
	Wander() error
	NMoves() int
}

type Location interface {
	Name() string
	Pieces() []Piece
	HasNeighbours() bool
	Neighbour(dir Direction) Location
	LinkToNeighbour(dir Direction, target Location)
	DeployPiece(p Piece)
}

type Direction string

func NewBoard() Board {
	panic("not implemented")
}

func NewLocation(name string) Location {
	panic("not implemented")
}
