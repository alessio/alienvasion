package board

import (
	"fmt"
	"strings"
)

func PrintBoard(b Board, indent int, outputAliens bool) {
	var w = b.(*worldMap)
	for _, location := range w.locations {
		printLocation(location, indent, outputAliens)
	}
}

func printLocation(loc Location, indent int, outputAliens bool) {
	var l = loc.(*location)
	s := strings.Repeat(" ", indent) + l.Name() + ":"
	for k, v := range l.neighbours {
		if v != nil {
			s += fmt.Sprintf(" %s=%s", string(k), v.Name())
		}
	}
	if outputAliens {
		s += " ["
		for _, v := range loc.Pieces() {
			s += string(v)
		}
		s += "]"
	}
	println(s)
}
