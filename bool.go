package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
)

var (
	_            fmt.Formatter          = (*Bool)(nil)
	_            json.Marshaler         = (*Bool)(nil)
	_            encoding.TextMarshaler = (*Bool)(nil)
	FormatBoolFn                        = func(s Bool, f fmt.State, c rune) {}
)

type Bool bool

func (s Bool) Format(f fmt.State, c rune) {
	FormatBoolFn(s, f, c)
}

func (s Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}

func (s Bool) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 's')
	return ss.b, nil
}
