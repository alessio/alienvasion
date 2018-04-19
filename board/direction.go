package board

const (
	North = "north"
	South = "south"
	West  = "west"
	East  = "east"
)

var oppositeDirectionsMap = map[Direction]Direction{
	North: South,
	South: North,
	West:  East,
	East:  West,
}

func (d Direction) String() string              { return string(d) }
func OppositeDirection(dir Direction) Direction { return oppositeDirectionsMap[dir] }
func (dir Direction) Validate() bool {
	_, ok := oppositeDirectionsMap[dir]
	return ok
}
