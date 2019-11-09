package main

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
