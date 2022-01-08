package scramble

import (
	"math"
)

func Scramble(id int32) int32 {
	var l1, l2, r1, r2 int32

	l1 = (id >> 16) & 65535
	r1 = id & 65535
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ int32(math.Round(float64(((1366.0*r1+150889)%714025))/714025.0*32767))
		l1 = l2
		r1 = r2
	}
	return ((r1 << 16) + l1)
}
