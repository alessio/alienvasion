package board

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeBoard(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		nlocations int
		wantErr    bool
	}{
		{"empty", "", 0, false},
		{"connected", `Baz north=Foo south=Bar east=Wxs west=Ming`, 5, false},
		{"inconsistent", "Baz north=Foo east=Bar\nWxs east=Bar\n", 5, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			b, err := DecodeBoard(reader)
			assert.Equal(t, tt.wantErr, (err != nil))
			if err == nil {
				assert.Equal(t, len(b.Locations()), tt.nlocations)
			}
		})
	}
}

func TestEncodeBoard(t *testing.T) {
	type args struct {
		b *Board
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			EncodeBoard(tt.args.b, writer)
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("EncodeBoard() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
