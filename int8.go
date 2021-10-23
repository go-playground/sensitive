package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	_            fmt.Formatter          = (*Int8)(nil)
	_            json.Marshaler         = (*Int8)(nil)
	_            encoding.TextMarshaler = (*Int8)(nil)
	FormatInt8Fn                        = func(s Int8, f fmt.State, c rune) {} //nolint:gochecknoglobals // By design.
)

type Int8 int8

func (s Int8) Format(f fmt.State, c rune) {
	FormatInt8Fn(s, f, c)
}

func (s Int8) MarshalJSON() ([]byte, error) {
	var ss State
	s.Format(&ss, 'v')
	if len(ss.b) == 0 {
		return json.Marshal(nil)
	}
	v, err := strconv.ParseInt(string(ss.b), base10, bits8)
	if err != nil {
		return nil, err
	}
	return json.Marshal(int8(v))
}

func (s Int8) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 'v')
	return ss.b, nil
}
