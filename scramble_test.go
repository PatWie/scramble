package scramble

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScramble32(t *testing.T) {

	cases := map[int32]int32{
		-10: -1270576520,
		-9:  -236348969,
		-8:  -1184061109,
		-7:  -25446276,
		-6:  -1507538963,
		-5:  -518858927,
		-4:  -1458116927,
		-3:  -532482573,
		-2:  -157973154,
		-1:  -1105881908,
		0:   1777613459,
		1:   561465857,
		2:   436885871,
		3:   576481439,
		4:   483424269,
		5:   1905133426,
		6:   971249312,
		7:   1926833684,
		8:   735327624,
		9:   1731020007,
		10:  792482838,
	}

	for id, want := range cases {
		got := Scramble32(id)
		assert.Equal(t, want, got)
		decrypt := Unscramble32(got)
		assert.Equal(t, id, decrypt)
	}

}

func TestScramble64(t *testing.T) {

	cases := map[int64]int64{
		-10: -54825524845095266,
		-9:  -3943936947812358396,
		-8:  -2033546936778524808,
		-7:  -2991612951414562926,
		-6:  -1208416983807528958,
		-5:  -918063562805821314,
		-4:  -4455032524137469812,
		-3:  -1734109410422572800,
		-2:  -2647536219340600970,
		-1:  -2271134734881302038,
		0:   1890057181288192996,
		1:   2652534423816930296,
		2:   667412736878761070,
		3:   1324215868541576419,
		4:   4491277083411598576,
		5:   174934417645116781,
		6:   4125364355537137114,
		7:   2253238117080874568,
		8:   2123696667436344939,
		9:   3586712234473600734,
		10:  1625706041578161985,
	}

	for id, want := range cases {
		got := Scramble64(id)
		assert.Equal(t, want, got)
		decrypt := Unscramble64(got)
		assert.Equal(t, id, decrypt)
	}
}
