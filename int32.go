package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
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
	return json.Marshal(nil)
}

func (s Int32) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 's')
	return ss.b, nil
}
