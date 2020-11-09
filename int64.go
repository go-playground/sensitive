package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	_             fmt.Formatter          = (*Int64)(nil)
	_             json.Marshaler         = (*Int64)(nil)
	_             encoding.TextMarshaler = (*Int64)(nil)
	FormatInt64Fn                        = func(s Int64, f fmt.State, c rune) {}
)

type Int64 int64

func (s Int64) Format(f fmt.State, c rune) {
	FormatInt64Fn(s, f, c)
}

func (s Int64) MarshalJSON() ([]byte, error) {
	var ss State
	s.Format(&ss, 'v')
	if len(ss.b) == 0 {
		return json.Marshal(nil)
	}
	v, err := strconv.ParseInt(string(ss.b), 10, 64)
	if err != nil {
		return nil, err
	}
	return json.Marshal(v)
}

func (s Int64) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 'v')
	return ss.b, nil
}
