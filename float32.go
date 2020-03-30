package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
)

var (
	_               fmt.Formatter          = (*Float32)(nil)
	_               json.Marshaler         = (*Float32)(nil)
	_               encoding.TextMarshaler = (*Float32)(nil)
	FormatFloat32Fn                        = func(s Float32, f fmt.State, c rune) {}
)

type Float32 float32

func (s Float32) Format(f fmt.State, c rune) {
	FormatFloat32Fn(s, f, c)
}

func (s Float32) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}

func (s Float32) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 's')
	return ss.b, nil
}
