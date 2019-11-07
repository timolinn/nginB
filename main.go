package main

import (
	"fmt"
)

// AttrMap is a map of attributes
// that belongs to an element
// Typically found within the ElementData
type AttrMap = map[string]string

func main() {
	fmt.Println(rune('A'), rune('Z'), rune('a'), rune('z'), rune('0'), rune('9'))
	fmt.Println(byte('A'), byte('Z'), byte('a'), byte('z'), byte('0'), byte('9'))
	bRune := rune(' ') // convert to rune for utf8 characters
	fmt.Println(bRune)
	isAplphaNum := ((bRune >= rune('a') && bRune <= rune('z')) ||
		(bRune >= rune('A') && bRune <= rune('Z')) ||
		(bRune >= rune('0') && bRune <= rune('9')))
	fmt.Println(isAplphaNum)
	s := "wwee23"
	fmt.Println([]byte('好'))
}
