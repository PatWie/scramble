package scramble

import (
	"math"
)

type Replacer32 func(int32) float64
type Replacer64 func(int64) float64

func DefaultReplace32(r1 int32) float64 {
	return float64(((1366.0*r1 + 150889) % 714025)) / 714025.0 * 32767
}

func DefaultReplace64(r1 int64) float64 {
	return float64(((1366.0*r1 + 150889) % 714025)) / 714025.0 * 32767.0 * 32767.0
}

func ScrambleWithFunction32(id int32, replacer Replacer32) int32 {
	var l1, l2, r1, r2 int32

	l1 = (id >> 16) & 65535
	r1 = id & 65535
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ int32(math.Round(replacer(r1)))
		l1 = l2
		r1 = r2
	}
	return ((r1 << 16) + l1)
}

func Scramble32(id int32) int32 {
	return ScrambleWithFunction32(id, DefaultReplace32)
}

func ScrambleWithFunction64(id int64, replacer Replacer64) int64 {
	var l1, l2, r1, r2 int64
	l1 = (id >> 32) & int64(4294967295)
	r1 = id & int64(4294967295)
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ int64(math.Round(replacer(r1)))
		l1 = l2
		r1 = r2
	}
	return ((r1 << 32) + l1)
}

func Scramble64(id int64) int64 {
	return ScrambleWithFunction64(id, DefaultReplace64)
}
