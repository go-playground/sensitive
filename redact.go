package sensitive

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/shopspring/decimal"
)

// Redact sets all Format<type>Fn to output visible, non-zero values:
//   Bool:    FALSE (in upper case, unlike usual bool)
//   Float*:  NaN
//   Int*:    math.MinInt* (MinInt32 for Int)
//   Uint*:   math.MaxUint* (MaxUint32 for Uint)
//   String:  "REDACTED"
//   Bytes:   0xDEFACE
//   Decimal: NaN
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
	FormatDecimalFn = func(s Decimal, f fmt.State, c rune) { Format(f, c, math.NaN()) }
}

// Disable protection of sensitive values.
//
// This is designed to be used only in tests to make it easier to compare
// got/want values. To make Disable actually works it's not enough to just
// call it, there are a couple of extra requirements to minimize a chance
// to get disabled sensitive in production:
//   - Current binary name should have ".test" suffix.
//   - Environment variable GO_TEST_DISABLE_SENSITIVE should not be empty.
//
// It is recommended to call it from TestMain or non-Parallel tests
// because it is not safe to call from simultaneous goroutines.
//
// Calling Redact after Disable will re-enable protection of sensitive values.
//
// As an extra protection it's recommended to add this into your main():
//   os.Unsetenv("GO_TEST_DISABLE_SENSITIVE")
func Disable() {
	if !strings.HasSuffix(os.Args[0], ".test") || os.Getenv("GO_TEST_DISABLE_SENSITIVE") == "" {
		return
	}
	FormatBoolFn = func(s Bool, f fmt.State, c rune) { Format(f, c, bool(s)) }
	FormatFloat32Fn = func(s Float32, f fmt.State, c rune) { Format(f, c, float32(s)) }
	FormatFloat64Fn = func(s Float64, f fmt.State, c rune) { Format(f, c, float64(s)) }
	FormatInt8Fn = func(s Int8, f fmt.State, c rune) { Format(f, c, int8(s)) }
	FormatInt16Fn = func(s Int16, f fmt.State, c rune) { Format(f, c, int16(s)) }
	FormatInt32Fn = func(s Int32, f fmt.State, c rune) { Format(f, c, int32(s)) }
	FormatInt64Fn = func(s Int64, f fmt.State, c rune) { Format(f, c, int64(s)) }
	FormatIntFn = func(s Int, f fmt.State, c rune) { Format(f, c, int(s)) }
	FormatUint8Fn = func(s Uint8, f fmt.State, c rune) { Format(f, c, uint8(s)) }
	FormatUint16Fn = func(s Uint16, f fmt.State, c rune) { Format(f, c, uint16(s)) }
	FormatUint32Fn = func(s Uint32, f fmt.State, c rune) { Format(f, c, uint32(s)) }
	FormatUint64Fn = func(s Uint64, f fmt.State, c rune) { Format(f, c, uint64(s)) }
	FormatUintFn = func(s Uint, f fmt.State, c rune) { Format(f, c, uint(s)) }
	FormatStringFn = func(s String, f fmt.State, c rune) { Format(f, c, string(s)) }
	FormatBytesFn = func(s Bytes, f fmt.State, c rune) { Format(f, c, []byte(s)) }
	FormatDecimalFn = func(s Decimal, f fmt.State, c rune) { Format(f, c, decimal.Decimal(s)) }
}
