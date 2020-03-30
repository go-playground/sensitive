package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
)

var (
	_            fmt.Formatter          = (*Uint)(nil)
	_            json.Marshaler         = (*Uint)(nil)
	_            encoding.TextMarshaler = (*Uint)(nil)
	FormatUintFn                        = func(s Uint, f fmt.State, c rune) {}
)

type Uint uint

func (s Uint) Format(f fmt.State, c rune) {
	FormatUintFn(s, f, c)
}

func (s Uint) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}

func (s Uint) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 's')
	return ss.b, nil
}
