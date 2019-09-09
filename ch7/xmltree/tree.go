package main

import (
	"encoding/xml"
	"fmt"
	"io"
)

// Node for node
type Node interface{}

// CharData for char
type CharData string

// Element for elem
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

// Parse generate a tree
func Parse(dec *xml.Decoder) (Node, error) {

	var stack []*Element
	var root *Element
	var top *Element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			return root, nil
		} else if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			var elem Element
			elem.Type = tok.Name
			for _, a := range tok.Attr {
				elem.Attr = append(elem.Attr, a)
			}
			if len(stack) == 0 {
				root = &elem
				stack = append(stack, &elem)
			} else {
				top = stack[len(stack)-1]

				top.Children = append(top.Children, elem)

				stack = append(stack, &elem)
			}

		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			var c CharData
			c = CharData(tok)
			top = stack[len(stack)-1]

			top.Children = append(top.Children, c)

		}

	}
}

// Print the tree
func Print(o io.Writer, root Node) {
	switch n := root.(type) {
	case *Element:
		fmt.Fprintf(o, "%s", n.Type)
		for _, child := range n.Children {
			Print(o, child)
		}
	case Element:
		fmt.Fprintf(o, "%s", n.Type)
		for _, child := range n.Children {
			Print(o, child)
		}
	case CharData:
		fmt.Fprintf(o, "%s", n)
	}

}
