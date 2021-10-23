package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	_              fmt.Formatter          = (*Uint32)(nil)
	_              json.Marshaler         = (*Uint32)(nil)
	_              encoding.TextMarshaler = (*Uint32)(nil)
	FormatUint32Fn                        = func(s Uint32, f fmt.State, c rune) {} //nolint:gochecknoglobals // By design.
)

type Uint32 uint32

func (s Uint32) Format(f fmt.State, c rune) {
	FormatUint32Fn(s, f, c)
}

func (s Uint32) MarshalJSON() ([]byte, error) {
	var ss State
	s.Format(&ss, 'v')
	if len(ss.b) == 0 {
		return json.Marshal(nil)
	}
	v, err := strconv.ParseUint(string(ss.b), base10, bits32)
	if err != nil {
		return nil, err
	}
	return json.Marshal(uint32(v))
}

func (s Uint32) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 'v')
	return ss.b, nil
}
