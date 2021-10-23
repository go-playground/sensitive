package sensitive_test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/shopspring/decimal"

	"github.com/powerman/sensitive"
)

func ExampleRedact() {
	var (
		vBool    sensitive.Bool    = true //nolint:stylecheck // False positive.
		vFloat32 sensitive.Float32 = 4.2
		vFloat64 sensitive.Float64 = 42.42
		vInt8    sensitive.Int8    = -42
		vInt16   sensitive.Int16   = -4242
		vInt32   sensitive.Int32   = -424242
		vInt64   sensitive.Int64   = -42424242
		vInt     sensitive.Int     = -42424242
		vUint8   sensitive.Uint8   = 42
		vUint16  sensitive.Uint16  = 4242
		vUint32  sensitive.Uint32  = 424242
		vUint64  sensitive.Uint64  = 42424242
		vUint    sensitive.Uint    = 42424242
		vString  sensitive.String  = "secret"
		vBytes   sensitive.Bytes   = []byte("secret")
		vDecimal                   = sensitive.Decimal(decimal.NewFromFloat(42.42))
		vs                         = []interface{}{
			vBool, vFloat32, vFloat64,
			vInt8, vInt16, vInt32, vInt64, vInt,
			vUint8, vUint16, vUint32, vUint64, vUint,
			vString, vBytes, vDecimal,
		}
		imap = map[interface{}]interface{}{
			vBool:    vBool,
			vFloat32: vFloat32,
			vFloat64: vFloat64,
			vInt8:    vInt8,
			vInt16:   vInt16,
			vInt32:   vInt32,
			vInt64:   vInt64,
			vInt:     vInt,
			vUint8:   vUint8,
			vUint16:  vUint16,
			vUint32:  vUint32,
			vUint64:  vUint64,
			vUint:    vUint,
			vString:  vString,
			// vBytes:   vBytes,
			vDecimal: vDecimal,
		}
		vmap = map[string]interface{}{
			"Bool":    vBool,
			"Float32": vFloat32,
			"Float64": vFloat64,
			"Int8":    vInt8,
			"Int16":   vInt16,
			"Int32":   vInt32,
			"Int64":   vInt64,
			"Int":     vInt,
			"Uint8":   vUint8,
			"Uint16":  vUint16,
			"Uint32":  vUint32,
			"Uint64":  vUint64,
			"Uint":    vUint,
			"String":  vString,
			"Bytes":   vBytes,
			"Decimal": vDecimal,
		}
		exported = struct {
			VBool    sensitive.Bool
			VFloat32 sensitive.Float32
			VFloat64 sensitive.Float64
			VInt8    sensitive.Int8
			VInt16   sensitive.Int16
			VInt32   sensitive.Int32
			VInt64   sensitive.Int64
			VInt     sensitive.Int
			VUint8   sensitive.Uint8
			VUint16  sensitive.Uint16
			VUint32  sensitive.Uint32
			VUint64  sensitive.Uint64
			VUint    sensitive.Uint
			VString  sensitive.String
			VBytes   sensitive.Bytes
			VDecimal sensitive.Decimal
		}{
			VBool:    vBool,
			VFloat32: vFloat32,
			VFloat64: vFloat64,
			VInt8:    vInt8,
			VInt16:   vInt16,
			VInt32:   vInt32,
			VInt64:   vInt64,
			VInt:     vInt,
			VUint8:   vUint8,
			VUint16:  vUint16,
			VUint32:  vUint32,
			VUint64:  vUint64,
			VUint:    vUint,
			VString:  vString,
			VBytes:   vBytes,
			VDecimal: vDecimal,
		}
		unexported = struct {
			vBool    sensitive.Bool
			vFloat32 sensitive.Float32
			vFloat64 sensitive.Float64
			vInt8    sensitive.Int8
			vInt16   sensitive.Int16
			vInt32   sensitive.Int32
			vInt64   sensitive.Int64
			vInt     sensitive.Int
			vUint8   sensitive.Uint8
			vUint16  sensitive.Uint16
			vUint32  sensitive.Uint32
			vUint64  sensitive.Uint64
			vUint    sensitive.Uint
			vString  sensitive.String
			vBytes   sensitive.Bytes
			// vDecimal Decimal
		}{
			vBool:    vBool,
			vFloat32: vFloat32,
			vFloat64: vFloat64,
			vInt8:    vInt8,
			vInt16:   vInt16,
			vInt32:   vInt32,
			vInt64:   vInt64,
			vInt:     vInt,
			vUint8:   vUint8,
			vUint16:  vUint16,
			vUint32:  vUint32,
			vUint64:  vUint64,
			vUint:    vUint,
			vString:  vString,
			vBytes:   vBytes,
			// vDecimal: vDecimal,
		}
	)

	// Outputs to stderr, not intercepted by testing package:
	// println: true +4.200000e+000 +4.242000e+001 -42 -4242 -424242 -42424242 -42424242 42 4242 424242 42424242 42424242 secret [6/6]0xc000216158
	println("println:",
		vBool, vFloat32, vFloat64,
		vInt8, vInt16, vInt32, vInt64, vInt,
		vUint8, vUint16, vUint32, vUint64, vUint,
		vString, vBytes,
		// vDecimal,
	)

	output := func() {
		fmt.Println(append(append([]interface{}{"fmt.Println(...):"}, vs...), "EOL")...)
		fmt.Printf("fmt.Printf: %t %e %E %c %b %o %x %d %c %b %O %X %U %q %X %v EOL\n", vs...)
		fmt.Printf("fmt.Printf(vs): %v\n", vs)
		fmt.Printf("fmt.Printf(imap): %v\n", imap)
		fmt.Printf("fmt.Printf(vmap): %v\n", vmap)
		fmt.Printf("fmt.Printf(exported): %v\n", exported)
		fmt.Printf("fmt.Printf(unexported): %v\n", unexported)
		json.NewEncoder(os.Stdout).Encode(vs)
		json.NewEncoder(os.Stdout).Encode(vmap)
		json.NewEncoder(os.Stdout).Encode(exported)
		xml.NewEncoder(os.Stdout).Encode(vs)
		fmt.Println()
	}
	output()
	sensitive.Redact()
	os.Unsetenv("GO_TEST_DISABLE_SENSITIVE")
	sensitive.Disable()
	output()
	os.Setenv("GO_TEST_DISABLE_SENSITIVE", "1")
	sensitive.Disable()
	output()
	// Output:
	// fmt.Println(...):                 EOL
	// fmt.Printf:                 EOL
	// fmt.Printf(vs): [               ]
	// fmt.Printf(imap): map[: : : : : : : : : : : : : : :]
	// fmt.Printf(vmap): map[Bool: Bytes: Decimal: Float32: Float64: Int: Int16: Int32: Int64: Int8: String: Uint: Uint16: Uint32: Uint64: Uint8:]
	// fmt.Printf(exported): {               }
	// fmt.Printf(unexported): {true 4.2 42.42 -42 -4242 -424242 -42424242 -42424242 42 4242 424242 42424242 42424242 secret [115 101 99 114 101 116]}
	// [null,null,null,null,null,null,null,null,null,null,null,null,null,"",null,null]
	// {"Bool":null,"Bytes":null,"Decimal":null,"Float32":null,"Float64":null,"Int":null,"Int16":null,"Int32":null,"Int64":null,"Int8":null,"String":"","Uint":null,"Uint16":null,"Uint32":null,"Uint64":null,"Uint8":null}
	// {"VBool":null,"VFloat32":null,"VFloat64":null,"VInt8":null,"VInt16":null,"VInt32":null,"VInt64":null,"VInt":null,"VUint8":null,"VUint16":null,"VUint32":null,"VUint64":null,"VUint":null,"VString":"","VBytes":null,"VDecimal":null}
	// <Bool></Bool><Float32></Float32><Float64></Float64><Int8></Int8><Int16></Int16><Int32></Int32><Int64></Int64><Int></Int><Uint8></Uint8><Uint16></Uint16><Uint32></Uint32><Uint64></Uint64><Uint></Uint><String></String><Bytes></Bytes><Decimal></Decimal>
	// fmt.Println(...): FALSE NaN NaN -128 -32768 -2147483648 -9223372036854775808 -2147483648 255 65535 4294967295 18446744073709551615 4294967295 REDACTED [222 250 206] NaN EOL
	// fmt.Printf: FALSE NaN NaN � -1000000000000000 -20000000000 -8000000000000000 -2147483648 ÿ 1111111111111111 0o37777777777 FFFFFFFFFFFFFFFF U+FFFFFFFF "REDACTED" DEFACE NaN EOL
	// fmt.Printf(vs): [FALSE NaN NaN -128 -32768 -2147483648 -9223372036854775808 -2147483648 255 65535 4294967295 18446744073709551615 4294967295 REDACTED [222 250 206] NaN]
	// fmt.Printf(imap): map[FALSE:FALSE NaN:NaN NaN:NaN -2147483648:-2147483648 -32768:-32768 -2147483648:-2147483648 -9223372036854775808:-9223372036854775808 -128:-128 REDACTED:REDACTED 4294967295:4294967295 65535:65535 4294967295:4294967295 18446744073709551615:18446744073709551615 255:255 NaN:NaN]
	// fmt.Printf(vmap): map[Bool:FALSE Bytes:[222 250 206] Decimal:NaN Float32:NaN Float64:NaN Int:-2147483648 Int16:-32768 Int32:-2147483648 Int64:-9223372036854775808 Int8:-128 String:REDACTED Uint:4294967295 Uint16:65535 Uint32:4294967295 Uint64:18446744073709551615 Uint8:255]
	// fmt.Printf(exported): {FALSE NaN NaN -128 -32768 -2147483648 -9223372036854775808 -2147483648 255 65535 4294967295 18446744073709551615 4294967295 REDACTED [222 250 206] NaN}
	// fmt.Printf(unexported): {true 4.2 42.42 -42 -4242 -424242 -42424242 -42424242 42 4242 424242 42424242 42424242 secret [115 101 99 114 101 116]}
	// <Bool>FALSE</Bool><Float32>NaN</Float32><Float64>NaN</Float64><Int8>-128</Int8><Int16>-32768</Int16><Int32>-2147483648</Int32><Int64>-9223372036854775808</Int64><Int>-2147483648</Int><Uint8>255</Uint8><Uint16>65535</Uint16><Uint32>4294967295</Uint32><Uint64>18446744073709551615</Uint64><Uint>4294967295</Uint><String>REDACTED</String><Bytes>DEFACE</Bytes><Decimal>NaN</Decimal>
	// fmt.Println(...): true 4.2 42.42 -42 -4242 -424242 -42424242 -42424242 42 4242 424242 42424242 42424242 secret [115 101 99 114 101 116] 42.42 EOL
	// fmt.Printf: true 4.200000e+00 4.242000E+01 � -1000010010010 -1474462 -28757b2 -42424242 * 1000010010010 0o1474462 28757B2 U+28757B2 "secret" 736563726574 42.42 EOL
	// fmt.Printf(vs): [true 4.2 42.42 -42 -4242 -424242 -42424242 -42424242 42 4242 424242 42424242 42424242 secret [115 101 99 114 101 116] 42.42]
	// fmt.Printf(imap): map[true:true 4.2:4.2 42.42:42.42 -42424242:-42424242 -4242:-4242 -424242:-424242 -42424242:-42424242 -42:-42 secret:secret 42424242:42424242 4242:4242 424242:424242 42424242:42424242 42:42 42.42:42.42]
	// fmt.Printf(vmap): map[Bool:true Bytes:[115 101 99 114 101 116] Decimal:42.42 Float32:4.2 Float64:42.42 Int:-42424242 Int16:-4242 Int32:-424242 Int64:-42424242 Int8:-42 String:secret Uint:42424242 Uint16:4242 Uint32:424242 Uint64:42424242 Uint8:42]
	// fmt.Printf(exported): {true 4.2 42.42 -42 -4242 -424242 -42424242 -42424242 42 4242 424242 42424242 42424242 secret [115 101 99 114 101 116] 42.42}
	// fmt.Printf(unexported): {true 4.2 42.42 -42 -4242 -424242 -42424242 -42424242 42 4242 424242 42424242 42424242 secret [115 101 99 114 101 116]}
	// [true,4.2,42.42,-42,-4242,-424242,-42424242,-42424242,42,4242,424242,42424242,42424242,"secret","c2VjcmV0",42.42]
	// {"Bool":true,"Bytes":"c2VjcmV0","Decimal":42.42,"Float32":4.2,"Float64":42.42,"Int":-42424242,"Int16":-4242,"Int32":-424242,"Int64":-42424242,"Int8":-42,"String":"secret","Uint":42424242,"Uint16":4242,"Uint32":424242,"Uint64":42424242,"Uint8":42}
	// {"VBool":true,"VFloat32":4.2,"VFloat64":42.42,"VInt8":-42,"VInt16":-4242,"VInt32":-424242,"VInt64":-42424242,"VInt":-42424242,"VUint8":42,"VUint16":4242,"VUint32":424242,"VUint64":42424242,"VUint":42424242,"VString":"secret","VBytes":"c2VjcmV0","VDecimal":42.42}
	// <Bool>true</Bool><Float32>4.2</Float32><Float64>42.42</Float64><Int8>-42</Int8><Int16>-4242</Int16><Int32>-424242</Int32><Int64>-42424242</Int64><Int>-42424242</Int><Uint8>42</Uint8><Uint16>4242</Uint16><Uint32>424242</Uint32><Uint64>42424242</Uint64><Uint>42424242</Uint><String>secret</String><Bytes>736563726574</Bytes><Decimal>42.42</Decimal>
}
