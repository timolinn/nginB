package main

import (
	"fmt"
)

func main() {
	parsed := Parse([]byte(`<html><body><p id="hw">Hello, world</p></body></html>`))
	fmt.Printf("%+v\n", parsed)
}
