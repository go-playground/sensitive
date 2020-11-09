package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
)

var (
	_             fmt.Formatter          = (*Bytes)(nil)
	_             json.Marshaler         = (*Bytes)(nil)
	_             encoding.TextMarshaler = (*Bytes)(nil)
	FormatBytesFn                        = func(s Bytes, f fmt.State, c rune) {}
)

type Bytes []byte

func (s Bytes) Format(f fmt.State, c rune) {
	FormatBytesFn(s, f, c)
}

func (s Bytes) MarshalJSON() ([]byte, error) {
	var ss State
	s.Format(&ss, 's')
	return json.Marshal(ss.b)
}

func (s Bytes) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 'X')
	return ss.b, nil
}
