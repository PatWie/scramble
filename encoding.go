package scramble

import (
	"math"
	"strings"
)

type (
	Encoding32 interface {
		Encode(int32) string
		Decode(string) int32
	}
	Encoding64 interface {
		Encode(int64) string
		Decode(string) int64
	}

	CharsetEncoder32 struct {
		Charset string
	}
	CharsetEncoder64 struct {
		Charset string
	}
)

var (
	UpperCase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerCase    = "abcdefghijklmnopqrstuvwxyz"
	AlphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// Encodes a integer int32 into a string using a given charset.
func (a CharsetEncoder32) Encode(value int32) string {
	base := int32(len(a.Charset))

	// Map negative part to positive side.
	if value < 0 {
		shift_multiplier := int32(float64(value)/float64(base) + 1)
		value += base * shift_multiplier
	}

	if value < base {
		return string(a.Charset[(value+base)%base])
	}

	output := ""
	for value > 0 {
		output = output + string(a.Charset[((value+base)%base)])
		value = value / base
	}
	return output
}

// Encodes a integer int64 into a string using a given charset.
func (a CharsetEncoder64) Encode(value int64) string {
	base := int64(len(a.Charset))

	// Map negative part to positive side.
	if value < 0 {
		shift_multiplier := int64(float64(value)/float64(base) + 1)
		value += base * shift_multiplier
	}

	if value < base {
		return string(a.Charset[(value+base)%base])
	}

	output := ""
	for value > 0 {
		offset := ((value + base) % base)
		letter := string(a.Charset[offset])
		output = output + letter
		value = value / base
	}
	return output
}

// Decodes a given string into an integer int32 using a given charset.
func (a CharsetEncoder32) Decode(value string) int32 {
	if len(value) == 1 {
		return int32(strings.Index(a.Charset, value))
	}

	base := int64(len(a.Charset))
	output := int32(0)
	for len(value) > 0 {
		letter := string(value[len(value)-1])
		offset := int32(strings.Index(a.Charset, letter))
		number := int32(math.Pow(float64(base), float64(len(value)-1))) * offset
		output += number
		value = value[:len(value)-1]
	}
	return output
}

// Decodes a given string into an integer int64 using a given charset.
func (a CharsetEncoder64) Decode(value string) int64 {
	if len(value) == 1 {
		return int64(strings.Index(a.Charset, value))
	}

	base := int64(len(a.Charset))
	output := int64(0)
	for len(value) > 0 {
		letter := string(value[len(value)-1])
		offset := int64(strings.Index(a.Charset, letter))
		number := int64(math.Pow(float64(base), float64(len(value)-1))) * offset
		output += number
		value = value[:len(value)-1]
	}
	return output
}
