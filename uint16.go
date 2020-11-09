package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	_              fmt.Formatter          = (*Uint16)(nil)
	_              json.Marshaler         = (*Uint16)(nil)
	_              encoding.TextMarshaler = (*Uint16)(nil)
	FormatUint16Fn                        = func(s Uint16, f fmt.State, c rune) {}
)

type Uint16 uint16

func (s Uint16) Format(f fmt.State, c rune) {
	FormatUint16Fn(s, f, c)
}

func (s Uint16) MarshalJSON() ([]byte, error) {
	var ss State
	s.Format(&ss, 'v')
	if len(ss.b) == 0 {
		return json.Marshal(nil)
	}
	v, err := strconv.ParseUint(string(ss.b), 10, 16)
	if err != nil {
		return nil, err
	}
	return json.Marshal(uint16(v))
}

func (s Uint16) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 'v')
	return ss.b, nil
}
