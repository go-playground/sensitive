package sensitive

import "fmt"

var (
	_ fmt.State = (*State)(nil)
)

type State struct {
	b []byte
}

func (s *State) Write(b []byte) (n int, err error) {
	s.b = append(s.b, b...)
	return len(b), nil
}

func (s State) Width() (wid int, ok bool) {
	return 0, false
}

func (s State) Precision() (prec int, ok bool) {
	return 0, false
}

func (s State) Flag(c int) bool {
	return false
}
