package board

import "fmt"

func PrintBoard(b Board) {
	var w = b.(*worldMap)
	for _, location := range w.locations {
		printLocation(location)
	}
	// for _, piece := range w.pieces {
	// 	s += fmt.Sprintf("%s\n", alien)
	// }
}

func printLocation(loc Location) {
	var l = loc.(*location)
	s := l.Name() + ":"
	for k, v := range l.neighbours {
		if v != nil {
			s += fmt.Sprintf(" %s=%s", string(k), v.Name())
		}
	}
	println(s)
}
