package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
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
	return json.Marshal(nil)
}

func (s Int64) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 's')
	return ss.b, nil
}
