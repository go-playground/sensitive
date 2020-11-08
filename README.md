## Package sensitive

<img align="right" src="https://raw.githubusercontent.com/go-playground/sensitive/master/logo.jpg">![Project status](https://img.shields.io/badge/version-0.0.1-green.svg)
[![Build Status](https://travis-ci.org/go-playground/sensitive.svg?branch=master)](https://travis-ci.org/go-playground/sensitive)
[![GoDoc](https://godoc.org/github.com/go-playground/sensitive?status.svg)](https://godoc.org/github.com/go-playground/sensitive)
![License](https://img.shields.io/dub/l/vibe-d.svg)

Package sensitive provides base types who's values should never be seen by the human eye, but still used for configuration.

What? Explain

Sometimes you have a variable, such as a password, passed into your program via arguments or ENV variables.
Some of these variables are very sensitive! and should not in any circumstance be loggged or sent via JSON, despite JSON's "-", which people may forget.
These variables, which are just typed primitive types, have their overridden `fmt.Formatter`, `encoding.MarshalText` & `json.Marshal` implementations.

As an added bonus using them as their base type eg. String => string, you have to explicitly cast the eg. string(s) This makes you think about what you're doing and why you casting it providing additional safelty.

Variables:
- `String` - The most useful
- `Bytes`
- `Bool`
- `Float32`
- `Float64`
- `Int`
- `Int8`
- `Int16`
- `Int32`
- `Int64`
- `Uint`
- `Uint8`
- `Uint16`
- `Uint32`
- `Uint64`

Example
-------
```go
// go run _examples/basic/main.go mypassword
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-playground/sensitive"
)

func main() {
	password := sensitive.String(os.Args[1])

	fmt.Printf("%s\n", password)
	fmt.Printf("%v\n", password)

	b, _ := json.Marshal(password)
	fmt.Println(string(b))

	var empty *sensitive.String
	b, _ = json.Marshal(empty)
	fmt.Println(string(b))

	// output:
	//
	//
	// ""
	// null
}
```

Custom Formatting
-----------------
```go
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-playground/sensitive"
)

func init() {
	// override default Formatter
	sensitive.FormatStringFn = func(s sensitive.String, f fmt.State, c rune) {
		switch c {
		default:
		        sensitive.Format(f, c, "redacted")
		case 'v':
		        sensitive.Format(f, c, string(s)[:4]+"*******")
		}
	}
}

func main() {
	password := sensitive.String(os.Args[1])

	fmt.Printf("%s\n", password)
	fmt.Printf("%v\n", password)

	b, _ := json.Marshal(password)
	fmt.Println(string(b))

	var empty *sensitive.String
	b, _ = json.Marshal(empty)
	fmt.Println(string(b))

	// output:
	// redacted
	// mypa*******
	// "redacted"
	// null
}
```

License
------
Distributed under MIT License, please see license file in code for more details.
