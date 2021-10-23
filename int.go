package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	_           fmt.Formatter          = (*Int)(nil)
	_           json.Marshaler         = (*Int)(nil)
	_           encoding.TextMarshaler = (*Int)(nil)
	FormatIntFn                        = func(s Int, f fmt.State, c rune) {} //nolint:gochecknoglobals // By design.
)

type Int int

func (s Int) Format(f fmt.State, c rune) {
	FormatIntFn(s, f, c)
}

func (s Int) MarshalJSON() ([]byte, error) {
	var ss State
	s.Format(&ss, 'v')
	if len(ss.b) == 0 {
		return json.Marshal(nil)
	}
	v, err := strconv.ParseInt(string(ss.b), base10, 0)
	if err != nil {
		return nil, err
	}
	return json.Marshal(int(v))
}

func (s Int) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 'v')
	return ss.b, nil
}
