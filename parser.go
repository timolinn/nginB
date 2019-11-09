package main

import (
	"errors"
	"strings"
	"unicode"
)

// AttrMap is a map of attributes
// that belongs to an element
// Typically found within the ElementData
type AttrMap = map[string]string

// Parser for html parsing
type Parser struct {
	pos   uint32
	input []byte
}

// NextChar returns the next char
func (p *Parser) NextChar() byte {
	return p.input[p.pos]
}

// StartsWith test whether input chars
// starts with s
func (p *Parser) StartsWith(s string) bool {
	return strings.HasPrefix(string(p.input[p.pos:]), s)
}

// EOF returns true when there's no string left to parse
func (p *Parser) EOF() bool {
	return int(p.pos) >= len(string(p.input))
}

// ConsumeChar returns the current character
// and increments the position
func (p *Parser) ConsumeChar() byte {
	currentChar := p.input[p.pos]
	p.pos++
	return currentChar
}

// ConsumeWhile consumes char until test returns false
func (p *Parser) ConsumeWhile(test func(byte) bool) []byte {
	var s []byte
	for !p.EOF() && test(p.NextChar()) {
		s = append(s, p.ConsumeChar())
	}
	return s
}

// ConsumeWhitespace consumes all the white space
func (p *Parser) ConsumeWhitespace() {
	p.ConsumeWhile(func(b byte) bool {
		return unicode.IsSpace(rune(b))
	})
}

// ParseTagName
func (p *Parser) ParseTagName() []byte {
	s := p.ConsumeWhile(func(b byte) bool {
		bRune := rune(b) // convert to rune for utf8 characters

		// Check if it is alphanumeric
		// regex might be a better way to achieve this
		// since we are checking this a character at a time
		// it should not be a problem
		return ((bRune >= rune('a') && bRune <= rune('z')) ||
			(bRune >= rune('A') && bRune <= rune('Z')) ||
			(bRune >= rune('0') && bRune <= rune('9')))
	})
	return s
}

// ParseNode parses a single node
func (p *Parser) ParseNode() Node {
	var n Node
	if p.NextChar() == '<' {
		n, _ = p.ParseElement()
	} else {
		n = p.ParseText()
	}
	return n
}

// ParseElement parses a single HTMLElement
func (p *Parser) ParseElement() (Node, error) {
	var n Node

	// start with <
	if p.ConsumeChar() != '<' {
		return n, errors.New("invalid HTMLElement: opening tag not found")
	}

	tagName := string(p.ParseTagName())
	attrs := p.ParseAttributes()
	if p.ConsumeChar() != '>' {
		return n, errors.New("invalid HTMLElement: invalid opening tag")
	} // End parsing

	// Parse contents
	children := p.ParseNodes()

	// parse element closing tag
	if p.ConsumeChar() != '<' {
		return n, errors.New("invalid HTMLElement: bad closing tag for <" + tagName + ">")
	}

	if p.ConsumeChar() != '/' {
		return n, errors.New("invalid HTMLElement: bad closing tag for <" + tagName + ">")
	}

	if string(p.ParseTagName()) != tagName {
		return n, errors.New("invalid HTMLElement: bad closing tag for <" + tagName + ">")
	}

	if p.ConsumeChar() != '>' {
		return n, errors.New("invalid HTMLElement: bad closing tag for <" + tagName + ">")
	}

	return Element(tagName, attrs, children), nil
}

// ParseText
func (p *Parser) ParseText() Node {
	var n Node
	t := string(p.ConsumeWhile(func(b byte) bool {
		return b != '<'
	}))
	n = Text(t)
	return n
}

// ParseAttr a single attr
func (p *Parser) ParseAttr() (string, string) {
	name := string(p.ParseTagName())
	if p.ConsumeChar() != '=' {
		panic("invalid html attr")
	}
	value := string(p.ParseAttrValue())
	return name, value
}

// ParseAttrValue
func (p *Parser) ParseAttrValue() []byte {
	openQuote := p.ConsumeChar()
	if openQuote != '"' || openQuote == '\'' {
		panic("Bad attr value")
	}
	value := p.ConsumeWhile(func(b byte) bool {
		return b != openQuote
	})
	if p.ConsumeChar() != openQuote {
		panic("Bad attr value")
	}
	return value
}

// ParseAttributes
func (p *Parser) ParseAttributes() AttrMap {
	attrs := make(AttrMap)
	for {
		p.ConsumeWhitespace()
		if p.NextChar() == '>' {
			break
		}
		name, value := p.ParseAttr()
		attrs[name] = value
	}

	return attrs
}

// ParseNodes parses all the nodes in a tree
func (p *Parser) ParseNodes() []Node {
	var nodes []Node
	for {
		p.ConsumeWhitespace()
		if p.EOF() || p.StartsWith("</") {
			break
		}
		nodes = append(nodes, p.ParseNode())
	}
	return nodes
}

// Parse parses string into html element
func Parse(source []byte) Node {
	p := Parser{pos: 0, input: source}
	nodes := p.ParseNodes()

	if len(nodes) == 1 {
		return nodes[0]
	}
	return Element("html", map[string]string{}, nodes)
}
