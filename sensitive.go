// Package sensitive provides base types who's values should never be seen
// by the human eye.
package sensitive

import "fmt"

const (
	base10 = 10
	bits8  = 8
	bits16 = 16
	bits32 = 32
	bits64 = 64
)

var _ fmt.State = (*State)(nil)

type State struct {
	b []byte
}

func (s *State) Write(b []byte) (n int, err error) {
	s.b = append(s.b, b...)
	return len(b), nil
}

func (State) Width() (wid int, ok bool) {
	return 0, false
}

func (State) Precision() (prec int, ok bool) {
	return 0, false
}

func (State) Flag(_ int) bool {
	return false
}
