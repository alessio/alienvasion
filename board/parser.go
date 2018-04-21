package board

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ParseBoard initializes a board configuration
// from an io.Reader's input.
func ParseBoard(reader io.Reader) (*Board, error) {
	lineNo := 0
	scanner := bufio.NewScanner(reader)
	b := NewBoard()
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		if len(tokens) < 2 {
			return nil, fmt.Errorf("line %d: invalid location (unlinked)", lineNo)
		}
		baseLocation := Location(tokens[0])
		b.AddLocation(baseLocation) // ignore error, it does not add if it exists
		for _, tok := range tokens[1:] {
			locDir := strings.Split(tok, "=")
			if len(locDir) != 2 {
				return nil, fmt.Errorf("line %d: couldn't parse token %q", lineNo, tok)
			}
			dir := Direction(locDir[0])
			target := Location(locDir[1])
			b.AddLocation(target)
			b.LinkLocations(baseLocation, target, dir)
		}
		lineNo++
	}
	return b, nil
}
