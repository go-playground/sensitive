package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
)

var (
	_             fmt.Formatter          = (*Int32)(nil)
	_             json.Marshaler         = (*Int32)(nil)
	_             encoding.TextMarshaler = (*Int32)(nil)
	FormatInt32Fn                        = func(s Int32, f fmt.State, c rune) {}
)

type Int32 int32

func (s Int32) Format(f fmt.State, c rune) {
	FormatInt32Fn(s, f, c)
}

func (s Int32) MarshalJSON() ([]byte, error) {
	var ss State
	s.Format(&ss, 'v')
	if len(ss.b) == 0 {
		return json.Marshal(nil)
	}
	v, err := strconv.ParseInt(string(ss.b), 10, 32)
	if err != nil {
		return nil, err
	}
	return json.Marshal(int32(v))
}

func (s Int32) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 'v')
	return ss.b, nil
}
