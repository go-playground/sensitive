package sensitive

import (
	"fmt"
	"strconv"
	"strings"
)

// Format outputs value accordingly to formatting options.
//
// It is useful in case you'll redefine some Format<type>Fn to output
// redacted value using formatting applied to original value.
//
//     sensitive.FormatStringFn = func(s sensitive.String, f fmt.State, c rune) {
//         sensitive.Format(f, c, "REDACTED")
//     }
//     sensitive.FormatBytesFn = func(s sensitive.Bytes, f fmt.State, c rune) {
//         sensitive.Format(f, c, []byte{0xDE, 0xFA, 0xCE})
//     }
func Format(f fmt.State, c rune, value interface{}) {
	const flags = "+-# 0"
	var format strings.Builder
	format.Grow(8)
	format.WriteRune('%')
	for _, c := range flags {
		if f.Flag(int(c)) {
			format.WriteRune(c)
		}
	}
	if wid, ok := f.Width(); ok {
		format.WriteString(strconv.Itoa(wid))
	}
	if prec, ok := f.Precision(); ok {
		format.WriteRune('.')
		format.WriteString(strconv.Itoa(prec))
	}
	format.WriteRune(c)
	fmt.Fprintf(f, format.String(), value)
}
