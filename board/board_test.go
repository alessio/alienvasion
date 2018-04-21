package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDirection_Opposite(t *testing.T) {
	tests := []struct {
		name      string
		d         Direction
		want      Direction
		wantPanic bool
	}{
		{"North", Direction(North), Direction(South), false},
		{"South", Direction(South), Direction(North), false},
		{"West", Direction(West), Direction(East), false},
		{"East", Direction(East), Direction(West), false},
		{"invalid", Direction("invalid"), Direction("invalid"), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() { tt.d.Opposite() })
				return
			}
			got := tt.d.Opposite()
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestNewBoard(t *testing.T) {
	got := NewBoard()
	assert.Nil(t, got.Locations())
	assert.Equal(t, got.Pieces(), make(map[Piece]Location))
}

func TestBoard_AddLocation(t *testing.T) {
	b := NewBoard()
	require.Nil(t, b.Locations())
	b.AddLocation("loc-00")
	assert.Equal(t, b.Locations(), []Location{Location("loc-00")})
	b.AddLocation("loc-00") // no-op
	assert.Equal(t, b.Locations(), []Location{Location("loc-00")})
	b.AddLocation("loc-01")
	b.AddLocation("loc-02")
	assert.Equal(t, len(b.Locations()), 3)
}

func TestBoard_LinkLocations(t *testing.T) {
	b := NewBoard()
	b.AddLocation(Location("0"))
	b.AddLocation(Location("1"))
	b.AddLocation(Location("2"))
	b.AddLocation(Location("3"))
	require.Equal(t, len(b.Locations()), 4)
	require.Equal(t, len(b.links), 4)
	b.LinkLocations(Location("0"), Location("1"), Direction("north"))
	b.LinkLocations(Location("0"), Location("2"), Direction("south"))
	b.LinkLocations(Location("2"), Location("3"), Direction("west"))
	assert.Equal(t, b.links[Location("0")][Direction("north")], Location("1"))
	assert.Equal(t, b.links[Location("1")][Direction("south")], Location("0"))
	assert.Equal(t, b.links[Location("0")][Direction("south")], Location("2"))
	assert.Equal(t, b.links[Location("2")][Direction("north")], Location("0"))
	assert.Equal(t, b.links[Location("2")][Direction("west")], Location("3"))
	assert.Equal(t, b.links[Location("3")][Direction("east")], Location("2"))
}

func TestBoard_Deploy(t *testing.T) {
	b := NewBoard()
	b.AddLocation(Location("0"))
	b.AddLocation(Location("1"))
	b.Deploy(Location("0"), Piece("a"))
	assert.Equal(t, len(b.Pieces()), 1)
	b.Deploy(Location("1"), Piece("b"))
	assert.Equal(t, len(b.Pieces()), 2)
}

func TestBoard_PieceNeighbourLocations(t *testing.T) {
	b := NewBoard()
	b.AddLocation(Location("0"))
	b.AddLocation(Location("1"))
	b.AddLocation(Location("2"))
	b.AddLocation(Location("3"))
	b.LinkLocations(Location("0"), Location("1"), Direction("north"))
	b.LinkLocations(Location("0"), Location("2"), Direction("south"))
	b.Deploy(Location("0"), Piece("a"))
	b.Deploy(Location("2"), Piece("b"))
	b.Deploy(Location("3"), Piece("c"))
	assert.Equal(t, len(b.PieceNeighbourLocations(Piece("a"))), 2)
	assert.Equal(t, len(b.PieceNeighbourLocations(Piece("b"))), 1)
	assert.Equal(t, len(b.PieceNeighbourLocations(Piece("c"))), 0)

}

func TestBoard_Wander(t *testing.T) {
	b := NewBoard()
	b.AddLocation(Location("0"))
	b.AddLocation(Location("1"))
	b.AddLocation(Location("2"))
	b.AddLocation(Location("3"))
	b.LinkLocations(Location("0"), Location("1"), Direction("north"))
	b.LinkLocations(Location("0"), Location("2"), Direction("south"))
	b.Deploy(Location("0"), Piece("a"))
	b.Deploy(Location("2"), Piece("b"))
	b.Deploy(Location("3"), Piece("c"))

	loc, err := b.Wander(Piece("c"))
	assert.Equal(t, loc, Location(""))
	assert.NotNil(t, err)
}

func TestBoard_move(t *testing.T) {
	b := NewBoard()
	b.AddLocation(Location("0"))
	b.AddLocation(Location("1"))
	b.AddLocation(Location("2"))
	b.LinkLocations(Location("0"), Location("1"), Direction("north"))
	b.Deploy(Location("0"), Piece("a"))
	require.Equal(t, b.pieces[Piece("a")], Location("0"))
	// Move back and forth
	b.move(Piece("a"), Location("1"))
	assert.Equal(t, b.pieces[Piece("a")], Location("1"))
	b.move(Piece("a"), Location("0"))
	assert.Equal(t, b.pieces[Piece("a")], Location("0"))
	// Test panic
	assert.Panics(t, func() { b.move(Piece("a"), Location("2")) })
}

func TestBoard_PiecesByLocation(t *testing.T) {
	b := NewBoard()
	b.AddLocation(Location("0"))
	b.AddLocation(Location("1"))
	b.AddLocation(Location("2"))
	b.Deploy(Location("0"), Piece("a"))
	b.Deploy(Location("0"), Piece("b"))
	b.Deploy(Location("1"), Piece("c"))
	piecesByLocation := b.PiecesByLocation()
	assert.Equal(t, len(piecesByLocation[Location("0")]), 2)
	assert.Equal(t, len(piecesByLocation[Location("1")]), 1)
	assert.Nil(t, piecesByLocation[Location("2")])
}

func TestBoard_Destroy(t *testing.T) {
	b := NewBoard()
	b.AddLocation(Location("0"))
	b.AddLocation(Location("1"))
	b.Deploy(Location("0"), Piece("a"))
	b.Deploy(Location("0"), Piece("b"))
	b.Deploy(Location("1"), Piece("c"))
	require.Equal(t, len(b.links), 2)
	require.Equal(t, len(b.pieces), 3)
	b.Destroy(Location("0"), []Piece{Piece("a"), Piece("b")})
	assert.Equal(t, len(b.links), 1)
	require.Equal(t, len(b.pieces), 1)
}
