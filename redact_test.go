package sensitive

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/shopspring/decimal"
)

func ExampleRedact() {
	var (
		vBool    Bool    = true
		vFloat32 Float32 = 4.2
		vFloat64 Float64 = 42.42
		vInt8    Int8    = -42
		vInt16   Int16   = -4242
		vInt32   Int32   = -424242
		vInt64   Int64   = -42424242
		vInt     Int     = -42424242
		vUint8   Uint8   = 42
		vUint16  Uint16  = 4242
		vUint32  Uint32  = 424242
		vUint64  Uint64  = 42424242
		vUint    Uint    = 42424242
		vString  String  = "secret"
		vBytes   Bytes   = []byte("secret")
		vDecimal Decimal = Decimal(decimal.NewFromFloat(42.42))
		vs               = []interface{}{
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
			VBool    Bool
			VFloat32 Float32
			VFloat64 Float64
			VInt8    Int8
			VInt16   Int16
			VInt32   Int32
			VInt64   Int64
			VInt     Int
			VUint8   Uint8
			VUint16  Uint16
			VUint32  Uint32
			VUint64  Uint64
			VUint    Uint
			VString  String
			VBytes   Bytes
			VDecimal Decimal
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
			vBool    Bool
			vFloat32 Float32
			vFloat64 Float64
			vInt8    Int8
			vInt16   Int16
			vInt32   Int32
			vInt64   Int64
			vInt     Int
			vUint8   Uint8
			vUint16  Uint16
			vUint32  Uint32
			vUint64  Uint64
			vUint    Uint
			vString  String
			vBytes   Bytes
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
		fmt.Printf("fmt.Printf: %t %e %E %c %b %o %x %d %c %b %O %X %U %q %X %E EOL\n", vs...)
		fmt.Printf("fmt.Printf(vs): %v\n", vs)
		fmt.Printf("fmt.Printf(imap): %v\n", imap)
		fmt.Printf("fmt.Printf(vmap): %v\n", vmap)
		fmt.Printf("fmt.Printf(exported): %v\n", exported)
		fmt.Printf("fmt.Printf(unexported): %v\n", unexported)
		json.NewEncoder(os.Stdout).Encode(vs)
		json.NewEncoder(os.Stdout).Encode(vmap)
		json.NewEncoder(os.Stdout).Encode(exported)
		json.NewEncoder(os.Stdout).Encode(unexported)
		xml.NewEncoder(os.Stdout).Encode(vs)
		fmt.Println()
	}
	output()
	Redact()
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
	// {}
	// <Bool></Bool><Float32></Float32><Float64></Float64><Int8></Int8><Int16></Int16><Int32></Int32><Int64></Int64><Int></Int><Uint8></Uint8><Uint16></Uint16><Uint32></Uint32><Uint64></Uint64><Uint></Uint><String></String><Bytes></Bytes><Decimal></Decimal>
	// fmt.Println(...): FALSE NaN NaN -128 -32768 -2147483648 -9223372036854775808 -2147483648 255 65535 4294967295 18446744073709551615 4294967295 REDACTED [222 250 206] NaN EOL
	// fmt.Printf: FALSE NaN NaN � -1000000000000000 -20000000000 -8000000000000000 -2147483648 ÿ 1111111111111111 0o37777777777 FFFFFFFFFFFFFFFF U+FFFFFFFF "REDACTED" DEFACE NaN EOL
	// fmt.Printf(vs): [FALSE NaN NaN -128 -32768 -2147483648 -9223372036854775808 -2147483648 255 65535 4294967295 18446744073709551615 4294967295 REDACTED [222 250 206] NaN]
	// fmt.Printf(imap): map[FALSE:FALSE NaN:NaN NaN:NaN -2147483648:-2147483648 -32768:-32768 -2147483648:-2147483648 -9223372036854775808:-9223372036854775808 -128:-128 REDACTED:REDACTED 4294967295:4294967295 65535:65535 4294967295:4294967295 18446744073709551615:18446744073709551615 255:255 NaN:NaN]
	// fmt.Printf(vmap): map[Bool:FALSE Bytes:[222 250 206] Decimal:NaN Float32:NaN Float64:NaN Int:-2147483648 Int16:-32768 Int32:-2147483648 Int64:-9223372036854775808 Int8:-128 String:REDACTED Uint:4294967295 Uint16:65535 Uint32:4294967295 Uint64:18446744073709551615 Uint8:255]
	// fmt.Printf(exported): {FALSE NaN NaN -128 -32768 -2147483648 -9223372036854775808 -2147483648 255 65535 4294967295 18446744073709551615 4294967295 REDACTED [222 250 206] NaN}
	// fmt.Printf(unexported): {true 4.2 42.42 -42 -4242 -424242 -42424242 -42424242 42 4242 424242 42424242 42424242 secret [115 101 99 114 101 116]}
	// {}
	// <Bool>FALSE</Bool><Float32>NaN</Float32><Float64>NaN</Float64><Int8>-128</Int8><Int16>-32768</Int16><Int32>-2147483648</Int32><Int64>-9223372036854775808</Int64><Int>-2147483648</Int><Uint8>255</Uint8><Uint16>65535</Uint16><Uint32>4294967295</Uint32><Uint64>18446744073709551615</Uint64><Uint>4294967295</Uint><String>REDACTED</String><Bytes>DEFACE</Bytes><Decimal>NaN</Decimal>
}
