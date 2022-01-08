package scramble

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadme(t *testing.T) {
	encoding := CharsetEncoder32{Charset: AlphaNumeric}

	internalID := int32(4)
	slug := encoding.Encode(Scramble32(internalID))
	assert.Equal(t, "NTySG", slug)
	assert.Equal(t, internalID, Unscramble32(encoding.Decode(slug)))

}

func TestReadmeInterface32(t *testing.T) {
	obscurer := NewObscurer32(CharsetEncoder32{Charset: AlphaNumeric})

	internalID := int32(4)

	assert.Equal(t, "NTySG", obscurer.Obscure(internalID))
	assert.Equal(t, internalID, obscurer.Unobscure("NTySG"))

}

func TestReadmeInterface64(t *testing.T) {
	obscurer := NewObscurer64(CharsetEncoder64{Charset: AlphaNumeric})

	internalID := int64(4)

	assert.Equal(t, "67LllQ9fWvf", obscurer.Obscure(internalID))
	assert.Equal(t, internalID, obscurer.Unobscure("67LllQ9fWvf"))

}
