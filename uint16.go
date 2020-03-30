package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
)

var (
	_              fmt.Formatter          = (*Uint16)(nil)
	_              json.Marshaler         = (*Uint16)(nil)
	_              encoding.TextMarshaler = (*Uint16)(nil)
	FormatUint16Fn                        = func(s Uint16, f fmt.State, c rune) {}
)

type Uint16 uint16

func (s Uint16) Format(f fmt.State, c rune) {
	FormatUint16Fn(s, f, c)
}

func (s Uint16) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}

func (s Uint16) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 's')
	return ss.b, nil
}
