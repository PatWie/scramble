package scramble

import (
	"math"
)

// TODO(patwie): A perfect candidate for generics?
type (
	// Settings for a int32 based linear congruence generator.
	LinearCongurenceParameters32 struct {
		a     float64
		c     int32
		mod   int32
		scale float64
	}

	// Settings for a int64 based linear congruence generator.
	LinearCongurenceParameters64 struct {
		a     float64
		c     int64
		mod   int64
		scale float64
	}

	LinearCongurencer32 interface {
		Transform(int32) float64
	}

	LinearCongurencer64 interface {
		Transform(int64) float64
	}
)

// A good setting of initial values (see C++ Boost) for int32.
var DefaultReplace32 = LinearCongurenceParameters32{
	a:     1366.0,
	c:     150889,
	mod:   714025,
	scale: 32767,
}

// Applies the linear congruence generator to an integer.
func (r LinearCongurenceParameters32) Transform(v int32) float64 {
	unbounded := (int32(r.a*float64(v)) + r.c) % r.mod
	return float64(unbounded) / float64(r.mod) * r.scale
}

// A good setting of initial values (see C++ Boost) for int64.
var DefaultReplace64 = LinearCongurenceParameters64{
	a:     1366.0,
	c:     150889,
	mod:   714025,
	scale: 32767 * 32767,
}

// Applies the linear congruence generator to an integer.
func (r LinearCongurenceParameters64) Transform(v int64) float64 {
	unbounded := (int64(r.a*float64(v)) + r.c) % r.mod
	return float64(unbounded) / float64(r.mod) * r.scale
}

// Runs 3 rounds of the Feistel cipher to encrypt an integer.
func FeistelNetwork32Crypt(start int32, f LinearCongurencer32) int32 {
	var l1, l2, r1, r2 int32

	l1 = (start >> 16) & 65535
	r1 = start & 65535
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ int32(math.Round(f.Transform(r1)))
		l1 = l2
		r1 = r2
	}
	return ((r1 << 16) + l1)
}

// Runs 3 rounds of the Feistel cipher to decrypt an integer.
func FeistelNetwork32Decrypt(start int32, f LinearCongurencer32) int32 {
	var l1, l2, r1, r2 int32

	r2 = (start >> 16) & 65535
	l2 = start & 65535
	for i := 0; i < 3; i++ {
		r1 = l2
		l1 = r2 ^ int32(math.Round(f.Transform(l2)))
		l2 = l1
		r2 = r1
	}
	return ((l2 << 16) + r2)
}

// Runs 3 rounds of the Feistel cipher to encrypt an integer.
func FeistelNetwork64Crypt(start int64, f LinearCongurencer64) int64 {
	var l1, l2, r1, r2 int64
	l1 = (start >> 32) & int64(4294967295)
	r1 = start & int64(4294967295)
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ int64(math.Round(f.Transform(r1)))
		l1 = l2
		r1 = r2
	}
	return ((r1 << 32) + l1)
}

// Runs 3 rounds of the Feistel cipher to decrypt an integer.
func FeistelNetwork64Decrypt(start int64, f LinearCongurencer64) int64 {
	var l1, l2, r1, r2 int64

	r2 = (start >> 32) & 4294967295
	l2 = start & 4294967295
	for i := 0; i < 3; i++ {
		r1 = l2
		l1 = r2 ^ int64(math.Round(f.Transform(l2)))
		l2 = l1
		r2 = r1
	}
	return ((l2 << 32) + r2)
}

// Scrambles an integer int32 in a bijective way.
func Scramble32(start int32) int32 {
	return FeistelNetwork32Crypt(start, DefaultReplace32)
}

// Inverse of Scramble32.
func Unscramble32(start int32) int32 {
	return FeistelNetwork32Decrypt(start, DefaultReplace32)
}

// Scrambles an integer int64 in a bijective way.
func Unscramble64(start int64) int64 {
	return FeistelNetwork64Decrypt(start, DefaultReplace64)
}

// Inverse of Scramble64.
func Scramble64(start int64) int64 {
	return FeistelNetwork64Crypt(start, DefaultReplace64)
}
