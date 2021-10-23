package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	_             fmt.Formatter          = (*Int16)(nil)
	_             json.Marshaler         = (*Int16)(nil)
	_             encoding.TextMarshaler = (*Int16)(nil)
	FormatInt16Fn                        = func(s Int16, f fmt.State, c rune) {} //nolint:gochecknoglobals // By design.
)

type Int16 int16

func (s Int16) Format(f fmt.State, c rune) {
	FormatInt16Fn(s, f, c)
}

func (s Int16) MarshalJSON() ([]byte, error) {
	var ss State
	s.Format(&ss, 'v')
	if len(ss.b) == 0 {
		return json.Marshal(nil)
	}
	v, err := strconv.ParseInt(string(ss.b), base10, bits16)
	if err != nil {
		return nil, err
	}
	return json.Marshal(int16(v))
}

func (s Int16) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 'v')
	return ss.b, nil
}
