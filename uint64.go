package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
)

var (
	_              fmt.Formatter          = (*Uint64)(nil)
	_              json.Marshaler         = (*Uint64)(nil)
	_              encoding.TextMarshaler = (*Uint64)(nil)
	FormatUint64Fn                        = func(s Uint64, f fmt.State, c rune) {}
)

type Uint64 uint64

func (s Uint64) Format(f fmt.State, c rune) {
	FormatUint64Fn(s, f, c)
}

func (s Uint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}

func (s Uint64) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 's')
	return ss.b, nil
}
