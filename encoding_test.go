package scramble

import (
	"math"
	"testing"

	"github.com/bits-and-blooms/bitset"

	"github.com/stretchr/testify/assert"
)

func TestEncoding32(t *testing.T) {

	cases := map[int32]string{
		-2: "Y",
		-1: "Z",
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

	encoding := CharsetEncoder32{Charset: UpperCase}
	for id, want := range cases {
		assert.Equal(t, want, encoding.Encode(id))
	}

}

func TestEncoding64(t *testing.T) {

	cases := map[int64]string{
		-2: "Y",
		-1: "Z",
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

	encoding := CharsetEncoder64{Charset: UpperCase}
	for id, want := range cases {
		assert.Equal(t, want, encoding.Encode(id))
	}

}

// Testing all would be too expensive.
func TestDecode32Encode32(t *testing.T) {
	encoding := CharsetEncoder32{Charset: UpperCase}
	for i := int32(0); i < math.MaxInt32-10000; i += 5000 {
		assert.Equal(t, i, encoding.Decode(encoding.Encode(i)))
	}
}

// Testing all would be too expensive.
func TestCollision32(t *testing.T) {
	var b bitset.BitSet
	for i := int32(0); i < math.MaxInt32-2000; i += 500 {
		pos := Scramble32(i)

		assert.GreaterOrEqual(t, pos, int32(0))
		assert.False(t, b.Test(uint(pos)))
		b.Set(uint(pos))
	}
}
