package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	_               fmt.Formatter          = (*Float64)(nil)
	_               json.Marshaler         = (*Float64)(nil)
	_               encoding.TextMarshaler = (*Float64)(nil)
	FormatFloat64Fn                        = func(s Float64, f fmt.State, c rune) {} //nolint:gochecknoglobals // By design.
)

type Float64 float64

func (s Float64) Format(f fmt.State, c rune) {
	FormatFloat64Fn(s, f, c)
}

func (s Float64) MarshalJSON() ([]byte, error) {
	var ss State
	s.Format(&ss, 'v')
	if len(ss.b) == 0 {
		return json.Marshal(nil)
	}
	v, err := strconv.ParseFloat(string(ss.b), bits64)
	if err != nil {
		return nil, err
	}
	return json.Marshal(v)
}

func (s Float64) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 'v')
	return ss.b, nil
}
