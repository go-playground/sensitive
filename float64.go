package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
)

var (
	_               fmt.Formatter          = (*Float64)(nil)
	_               json.Marshaler         = (*Float64)(nil)
	_               encoding.TextMarshaler = (*Float64)(nil)
	FormatFloat64Fn                        = func(s Float64, f fmt.State, c rune) {}
)

type Float64 float64

func (s Float64) Format(f fmt.State, c rune) {
	FormatFloat64Fn(s, f, c)
}

func (s Float64) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}

func (s Float64) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 's')
	return ss.b, nil
}
