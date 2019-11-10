package main

import (
	"fmt"

	"github.com/timolinn/html-parser"
)

func main() {
	parsed := html.Parse([]byte(`<html><body><p id="hw">Hello, world</p></body></html>`))
	fmt.Printf("%+v\n", parsed)
}
