package scramble

type (
	Obscurer32 interface {
		Obscure(int32) string
		Unobscure(string) int32
	}

	Obscurer64 interface {
		Obscure(int64) string
		Unobscure(string) int64
	}

	// Obscurer32 scrambles 32 bit integers.
	ObscureObscurer32 struct {
		Encoding Encoding32
	}
	// Obscurer32 scrambles 32 bit integers.
	ObscureObscurer64 struct {
		Encoding Encoding64
	}
)

func (s ObscureObscurer32) Obscure(start int32) string {
	return s.Encoding.Encode(Scramble32(start))
}

func (s ObscureObscurer32) Unobscure(start string) int32 {
	return Unscramble32(s.Encoding.Decode(start))
}

func NewObscurer32(encoding Encoding32) Obscurer32 {
	return ObscureObscurer32{Encoding: encoding}
}

func (s ObscureObscurer64) Obscure(start int64) string {
	return s.Encoding.Encode(Scramble64(start))
}
func (s ObscureObscurer64) Unobscure(start string) int64 {
	return Unscramble64(s.Encoding.Decode(start))
}
func NewObscurer64(encoding Encoding64) Obscurer64 {
	return ObscureObscurer64{Encoding: encoding}
}
