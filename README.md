# Go package with base types protected from the human eye

[![Go Reference](https://pkg.go.dev/badge/github.com/powerman/sensitive.svg)](https://pkg.go.dev/github.com/powerman/sensitive)
[![CI/CD](https://github.com/powerman/sensitive/workflows/CI/CD/badge.svg?event=push)](https://github.com/powerman/sensitive/actions?query=workflow%3ACI%2FCD)
[![Coverage Status](https://coveralls.io/repos/github/powerman/sensitive/badge.svg?branch=master)](https://coveralls.io/github/powerman/sensitive?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/powerman/sensitive)](https://goreportcard.com/report/github.com/powerman/sensitive)
[![Release](https://img.shields.io/github/v/release/powerman/sensitive)](https://github.com/powerman/sensitive/releases/latest)

**NOTE:** This projects starts as a fork, but as upstream have no activity
since initial commit new features are added here (see
[Releases](https://github.com/powerman/sensitive/releases)) and you can
consider this fork a maintained version of the upstream repo.

Package sensitive provides base types who's values should never be seen by the human eye, but still used for configuration.

What? Explain

Sometimes you have a variable, such as a password, passed into your program via arguments or ENV variables.
Some of these variables are very sensitive! and should not in any circumstance be loggged or sent via JSON, despite JSON's "-", which people may forget.
These variables, which are just typed primitive types, have their overridden `fmt.Formatter`, `encoding.MarshalText` & `json.Marshal` implementations.

As an added bonus using them as their base type eg. String => string, you have to explicitly cast the eg. string(s) This makes you think about what you're doing and why you casting it providing additional safelty.

Variables:
- `Bool`
- `Bytes`
- `Decimal` (for https://github.com/shopspring/decimal)
- `Float32`
- `Float64`
- `Int`
- `Int8`
- `Int16`
- `Int32`
- `Int64`
- `String` - The most useful
- `Uint`
- `Uint8`
- `Uint16`
- `Uint32`
- `Uint64`

## Examples

### Basic

```go
// go run _examples/basic/main.go mypassword
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/powerman/sensitive"
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

### Custom Formatting

```go
// go run _examples/custom/main.go mypassword
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/powerman/sensitive"
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
	// "mypa*******"
	// null
}
```
