package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/go-playground/sensitive"
)

func init() {
	// override default Formatter
	sensitive.FormatStringFn = func(s sensitive.String, f fmt.State, c rune) {
		switch c {
		case 's':
			_, _ = io.WriteString(f, "redacted")
		case 'v':
			_, _ = io.WriteString(f, string(s)[:4]+"*******")
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
