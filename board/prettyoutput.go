package board

import (
	"fmt"
)

func EncodeBoard(b *Board) string {
	var s string
	for baselocation, dirmap := range b.links {
		s += fmt.Sprintf("%s:", baselocation)
		for dir, target := range dirmap {
			s += fmt.Sprintf(" %s=%s", dir, target)
		}
		s += "\n"
	}
	return s
}
