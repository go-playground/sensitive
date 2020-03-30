package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
)

var (
	_             fmt.Formatter          = (*Uint8)(nil)
	_             json.Marshaler         = (*Uint8)(nil)
	_             encoding.TextMarshaler = (*Uint8)(nil)
	FormatUint8Fn                        = func(s Uint8, f fmt.State, c rune) {}
)

type Uint8 uint8

func (s Uint8) Format(f fmt.State, c rune) {
	FormatUint8Fn(s, f, c)
}

func (s Uint8) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}

func (s Uint8) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 's')
	return ss.b, nil
}
