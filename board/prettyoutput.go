package board

import (
	"fmt"
	"strings"
)

func PrintBoard(b Board, indent int) {
	var w = b.(*worldMap)
	for _, location := range w.locations {
		printLocation(location, indent)
	}
}

func printLocation(loc Location, indent int) {
	var l = loc.(*location)
	s := strings.Repeat(" ", indent) + l.Name() + ":"
	for k, v := range l.neighbours {
		if v != nil {
			s += fmt.Sprintf(" %s=%s", string(k), v.Name())
		}
	}
	fmt.Println(s)
}
