package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
)

var (
	_              fmt.Formatter          = (*Uint32)(nil)
	_              json.Marshaler         = (*Uint32)(nil)
	_              encoding.TextMarshaler = (*Uint32)(nil)
	FormatUint32Fn                        = func(s Uint32, f fmt.State, c rune) {}
)

type Uint32 uint32

func (s Uint32) Format(f fmt.State, c rune) {
	FormatUint32Fn(s, f, c)
}

func (s Uint32) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}

func (s Uint32) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 's')
	return ss.b, nil
}
