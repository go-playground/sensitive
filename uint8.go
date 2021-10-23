package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	_             fmt.Formatter          = (*Uint8)(nil)
	_             json.Marshaler         = (*Uint8)(nil)
	_             encoding.TextMarshaler = (*Uint8)(nil)
	FormatUint8Fn                        = func(s Uint8, f fmt.State, c rune) {} //nolint:gochecknoglobals // By design.
)

type Uint8 uint8

func (s Uint8) Format(f fmt.State, c rune) {
	FormatUint8Fn(s, f, c)
}

func (s Uint8) MarshalJSON() ([]byte, error) {
	var ss State
	s.Format(&ss, 'v')
	if len(ss.b) == 0 {
		return json.Marshal(nil)
	}
	v, err := strconv.ParseUint(string(ss.b), base10, bits8)
	if err != nil {
		return nil, err
	}
	return json.Marshal(uint8(v))
}

func (s Uint8) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 'v')
	return ss.b, nil
}
