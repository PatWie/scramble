package scramble

type (
	Encoder interface {
		Encode32(int64) string
		Encode64(int64) string
	}

	AlphabethEncoder struct {
		Alphabeth string
	}
)

var (
	UpperCaseEncoder = AlphabethEncoder{
		Alphabeth: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}
	LowerCaseEncoder = AlphabethEncoder{
		Alphabeth: "abcdefghijklmnopqrstuvwxyz",
	}

	AlphaNumericEncoder = AlphabethEncoder{
		Alphabeth: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	}
)

func (a AlphabethEncoder) Encode32(value int32) string {
	base := int32(len(a.Alphabeth))
	if value < 0 {
		value *= -1
	}
	if value < base {
		return string(a.Alphabeth[value])
	}

	output := ""

	for value > 0 {
		output = output + string(a.Alphabeth[(value%base)])
		value = value / base
	}
	return output
}

func (a AlphabethEncoder) Encode64(value int64) string {
	base := int64(len(a.Alphabeth))
	if value < 0 {
		value *= -1
	}
	if value < base {
		return string(a.Alphabeth[value])
	}

	output := ""

	for value > 0 {
		output = output + string(a.Alphabeth[(value%base)])
		value = value / base
	}
	return output
}
