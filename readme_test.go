package scramble

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadme(t *testing.T) {
	encoding := AlphabethEncoder32{Alphabeth: AlphaNumeric}

	internalID := int32(4)
	slug := encoding.Encode(Scramble32(internalID))
	assert.Equal(t, "NTySG", slug)
	assert.Equal(t, internalID, Unscramble32(encoding.Decode(slug)))

}