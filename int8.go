package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
)

var (
	_            fmt.Formatter          = (*Int8)(nil)
	_            json.Marshaler         = (*Int8)(nil)
	_            encoding.TextMarshaler = (*Int8)(nil)
	FormatInt8Fn                        = func(s Int8, f fmt.State, c rune) {}
)

type Int8 int8

func (s Int8) Format(f fmt.State, c rune) {
	FormatInt8Fn(s, f, c)
}

func (s Int8) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}

func (s Int8) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 's')
	return ss.b, nil
}
