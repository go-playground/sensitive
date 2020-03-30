package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
)

var (
	_             fmt.Formatter          = (*Int16)(nil)
	_             json.Marshaler         = (*Int16)(nil)
	_             encoding.TextMarshaler = (*Int16)(nil)
	FormatInt16Fn                        = func(s Int16, f fmt.State, c rune) {}
)

type Int16 int16

func (s Int16) Format(f fmt.State, c rune) {
	FormatInt16Fn(s, f, c)
}

func (s Int16) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}

func (s Int16) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 's')
	return ss.b, nil
}
