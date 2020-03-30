package sensitive

import (
	"encoding"
	"encoding/json"
	"fmt"
)

var (
	_           fmt.Formatter          = (*Int)(nil)
	_           json.Marshaler         = (*Int)(nil)
	_           encoding.TextMarshaler = (*Int)(nil)
	FormatIntFn                        = func(s Int, f fmt.State, c rune) {}
)

type Int int

func (s Int) Format(f fmt.State, c rune) {
	FormatIntFn(s, f, c)
}

func (s Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(nil)
}

func (s Int) MarshalText() (text []byte, err error) {
	var ss State
	s.Format(&ss, 's')
	return ss.b, nil
}
