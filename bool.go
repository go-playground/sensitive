package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	_            fmt.Formatter          = (*Bool)(nil)
	_            json.Marshaler         = (*Bool)(nil)
	_            encoding.TextMarshaler = (*Bool)(nil)
	FormatBoolFn                        = func(s Bool, f fmt.State, c rune) {} //nolint:gochecknoglobals // By design.
)

type Bool bool

func (s Bool) Format(f fmt.State, c rune) {
	FormatBoolFn(s, f, c)
}

func (s Bool) MarshalJSON() ([]byte, error) {
	var ss State
	s.Format(&ss, 'v')
	if len(ss.b) == 0 {
		return json.Marshal(nil)
	}
	v, err := strconv.ParseBool(string(ss.b))
	if err != nil {
		return nil, err
	}
	return json.Marshal(v)
}

func (s Bool) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 'v')
	return ss.b, nil
}
