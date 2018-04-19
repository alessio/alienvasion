package board

// Board defines the
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

type Location interface {
	Name() string
	Pieces() []Piece
	ReachableNeighbours() []Location
	Neighbour(dir Direction) Location
	LinkToNeighbour(dir Direction, target Location)
	AddPiece(p Piece)
	RemovePiece(p Piece)
}

type Direction string
