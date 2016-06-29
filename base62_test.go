package base62

import (
	"bytes"
	"github.com/growsumo/base62"
	"testing"
)

var cases = map[int][]byte{
	0:          []byte("0"),
	1:          []byte("1"),
	100:        []byte("1C"),
	999:        []byte("g7"),
	9999999999: []byte("aUKYOz"),
}

func TestEncode(t *testing.T) {
	for input, expected := range cases {
		output := base62.Encode(input)
		if !bytes.Equal(expected, output) {
			t.Errorf("Encode(%d), Got: %d,  Expected: %d", input, output, expected)
		}
	}
}

func TestDecode(t *testing.T) {
	for expected, input := range cases {
		output := base62.Decode(input)
		if expected != output {
			t.Errorf("Decode(%s) => Got: %d,  Expected: %d", input, output, expected)
		}
	}
}
