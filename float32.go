package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	_               fmt.Formatter          = (*Float32)(nil)
	_               json.Marshaler         = (*Float32)(nil)
	_               encoding.TextMarshaler = (*Float32)(nil)
	FormatFloat32Fn                        = func(s Float32, f fmt.State, c rune) {} //nolint:gochecknoglobals // By design.
)

type Float32 float32

func (s Float32) Format(f fmt.State, c rune) {
	FormatFloat32Fn(s, f, c)
}

func (s Float32) MarshalJSON() ([]byte, error) {
	var ss State
	s.Format(&ss, 'v')
	if len(ss.b) == 0 {
		return json.Marshal(nil)
	}
	v, err := strconv.ParseFloat(string(ss.b), bits32)
	if err != nil {
		return nil, err
	}
	return json.Marshal(float32(v))
}

func (s Float32) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 'v')
	return ss.b, nil
}
