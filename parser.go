package main

import (
	"strings"
	"unicode"
)

type NodeType int

const (
	TextNode NodeType = iota
	HTMLElement
)

type Node struct {
	Children []Node
	NodeType NodeType
}

func Text(data string) Node {
	return Node{
		Children: []Node{},
		NodeType: TextNode,
	}
}

func Element(name string, attrs AttrMap, children []Node) Node {
	return Node{
		Children: children,
		NodeType: HTMLElement,
	}
}

type Parser struct {
	pos   uint32
	input []byte
}

func (p *Parser) NextChar() byte {
	return p.input[p.pos]
}

func (p *Parser) StartsWith(s string) bool {
	return strings.HasPrefix(string(p.input[p.pos:]), s)
}

func (p *Parser) Eof() bool {
	return int(p.pos) > len(p.input)
}

func (p *Parser) ConsumeChar() byte {
	currentChar := p.input[p.pos]
	p.pos++
	return currentChar
}

func (p *Parser) ConsumeWhile(test func(byte) bool) []byte {
	var s []byte
	for !p.Eof() && test(p.NextChar()) {
		s = append(s, p.ConsumeChar())
	}
	return s
}

func (p *Parser) ConsumeWhitespace() {
	p.ConsumeWhile(func(b byte) bool {
		return unicode.IsSpace(rune(b))
	})
}

func (p *Parser) ParseTagName() []byte {
	s := p.ConsumeWhile(func(b byte) bool {
		bRune := rune(b) // convert to rune for utf8 characters

		// regex might be a better way to achieve this
		// since we are checking this a character at a time
		// it should not be a problem
		return ((bRune >= rune('a') && bRune <= rune('z')) ||
			(bRune >= rune('A') && bRune <= rune('Z')) ||
			(bRune >= rune('0') && bRune <= rune('9')))
	})
	return s
}
