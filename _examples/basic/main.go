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
