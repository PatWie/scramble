package scramble

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode32(t *testing.T) {

	cases := map[int32]string{
		0:  "A",
		1:  "B",
		2:  "C",
		3:  "D",
		4:  "E",
		25: "Z",
		26: "AB",
		27: "BB",
		28: "CB",
		52: "AC",
		53: "BC",
	}

	encoder := AlphabethEncoder{
		Alphabeth: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}

	for id, want := range cases {
		assert.Equal(t, want, encoder.Encode32(id))
		assert.Equal(t, want, encoder.Encode32(-id))
	}

}

func TestEncode64(t *testing.T) {

	cases := map[int64]string{
		0:  "A",
		1:  "B",
		2:  "C",
		3:  "D",
		4:  "E",
		25: "Z",
		26: "AB",
		27: "BB",
		28: "CB",
		52: "AC",
		53: "BC",
	}

	encoder := AlphabethEncoder{
		Alphabeth: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}

	for id, want := range cases {
		assert.Equal(t, want, encoder.Encode64(id))
		assert.Equal(t, want, encoder.Encode64(-id))
	}

}
