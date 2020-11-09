package sensitive

import (
	"fmt"
	"math"
)

// Redact sets all Format<type>Fn to output visible, non-zero values:
//   Bool:    FALSE (in upper case, unlike usual bool)
//   Float*:  NaN
//   Int*:    math.MinInt* (MinInt32 for Int)
//   Uint*:   math.MaxUint* (MaxUint32 for Uint)
//   String:  "REDACTED"
//   Bytes:   0xDEFACE
func Redact() {
	FormatBoolFn = func(s Bool, f fmt.State, c rune) { Format(f, 's', "FALSE") }
	FormatFloat32Fn = func(s Float32, f fmt.State, c rune) { Format(f, c, float32(math.NaN())) }
	FormatFloat64Fn = func(s Float64, f fmt.State, c rune) { Format(f, c, math.NaN()) }
	FormatInt8Fn = func(s Int8, f fmt.State, c rune) { Format(f, c, int8(math.MinInt8)) }
	FormatInt16Fn = func(s Int16, f fmt.State, c rune) { Format(f, c, int16(math.MinInt16)) }
	FormatInt32Fn = func(s Int32, f fmt.State, c rune) { Format(f, c, int32(math.MinInt32)) }
	FormatInt64Fn = func(s Int64, f fmt.State, c rune) { Format(f, c, int64(math.MinInt64)) }
	FormatIntFn = func(s Int, f fmt.State, c rune) { Format(f, c, int(math.MinInt32)) }
	FormatUint8Fn = func(s Uint8, f fmt.State, c rune) { Format(f, c, uint8(math.MaxUint8)) }
	FormatUint16Fn = func(s Uint16, f fmt.State, c rune) { Format(f, c, uint16(math.MaxUint16)) }
	FormatUint32Fn = func(s Uint32, f fmt.State, c rune) { Format(f, c, uint32(math.MaxUint32)) }
	FormatUint64Fn = func(s Uint64, f fmt.State, c rune) { Format(f, c, uint64(math.MaxUint64)) }
	FormatUintFn = func(s Uint, f fmt.State, c rune) { Format(f, c, uint(math.MaxUint32)) }
	FormatStringFn = func(s String, f fmt.State, c rune) { Format(f, c, "REDACTED") }
	FormatBytesFn = func(s Bytes, f fmt.State, c rune) { Format(f, c, []byte{0xDE, 0xFA, 0xCE}) }
}
