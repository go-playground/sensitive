package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	_              fmt.Formatter          = (*Uint64)(nil)
	_              json.Marshaler         = (*Uint64)(nil)
	_              encoding.TextMarshaler = (*Uint64)(nil)
	FormatUint64Fn                        = func(s Uint64, f fmt.State, c rune) {} //nolint:gochecknoglobals // By design.
)

type Uint64 uint64

func (s Uint64) Format(f fmt.State, c rune) {
	FormatUint64Fn(s, f, c)
}

func (s Uint64) MarshalJSON() ([]byte, error) {
	var ss State
	s.Format(&ss, 'v')
	if len(ss.b) == 0 {
		return json.Marshal(nil)
	}
	v, err := strconv.ParseUint(string(ss.b), base10, bits64)
	if err != nil {
		return nil, err
	}
	return json.Marshal(v)
}

func (s Uint64) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 'v')
	return ss.b, nil
}
