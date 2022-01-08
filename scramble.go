package scramble

import (
	"math"
)

// TODO(patwie): A perfect candidate for generics?
type (
	LinearCongurenceParameters32 struct {
		a     float64
		c     int32
		mod   int32
		scale float64
	}

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

var DefaultReplace32 = LinearCongurenceParameters32{
	a:     1366.0,
	c:     150889,
	mod:   714025,
	scale: 32767,
}

func (r LinearCongurenceParameters32) Transform(v int32) float64 {
	unbounded := (int32(r.a*float64(v)) + r.c) % r.mod
	return float64(unbounded) / float64(r.mod) * r.scale
}

var DefaultReplace64 = LinearCongurenceParameters64{
	a:     1366.0,
	c:     150889,
	mod:   714025,
	scale: 32767 * 32767,
}

func (r LinearCongurenceParameters64) Transform(v int64) float64 {
	unbounded := (int64(r.a*float64(v)) + r.c) % r.mod
	return float64(unbounded) / float64(r.mod) * r.scale
}

func FeistelNetwork32(start int32, f LinearCongurencer32) int32 {
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

func Scramble32(start int32) int32 {
	return FeistelNetwork32(start, DefaultReplace32)
}

func FeistelNetwork64(start int64, f LinearCongurencer64) int64 {
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

func Scramble64(start int64) int64 {
	return FeistelNetwork64(start, DefaultReplace64)
}
