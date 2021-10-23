package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	_            fmt.Formatter          = (*Uint)(nil)
	_            json.Marshaler         = (*Uint)(nil)
	_            encoding.TextMarshaler = (*Uint)(nil)
	FormatUintFn                        = func(s Uint, f fmt.State, c rune) {} //nolint:gochecknoglobals // By design.
)

type Uint uint

func (s Uint) Format(f fmt.State, c rune) {
	FormatUintFn(s, f, c)
}

func (s Uint) MarshalJSON() ([]byte, error) {
	var ss State
	s.Format(&ss, 'v')
	if len(ss.b) == 0 {
		return json.Marshal(nil)
	}
	v, err := strconv.ParseUint(string(ss.b), base10, 0)
	if err != nil {
		return nil, err
	}
	return json.Marshal(uint(v))
}

func (s Uint) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 'v')
	return ss.b, nil
}
