package board

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ParseBoard initializes a board configuration
// from an io.Reader's input.
func DecodeBoard(reader io.Reader) (*Board, error) {
	scanner := bufio.NewScanner(reader)
	b := NewBoard()
	for lineNo := 0; scanner.Scan(); lineNo++ {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue // skip comments
		}
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
			if err := b.LinkLocations(baseLocation, target, dir); err != nil {
				return nil, fmt.Errorf("line %d: couldn't link %s and %s: %v", lineNo, baseLocation, target, err)
			}
		}
	}
	return b, nil
}

func EncodeBoard(b *Board, writer io.Writer) {
	var s string
	for baselocation, dirmap := range b.links {
		s += fmt.Sprintf("%s:", baselocation)
		for dir, target := range dirmap {
			s += fmt.Sprintf(" %s=%s", dir, target)
		}
		s += "\n"
	}
	fmt.Fprint(writer, s)
}
