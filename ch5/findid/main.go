// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 134. Exercise 5.7

// Outline prints the outline of an HTML document tree.
package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var f = flag.String("f", "", "html file name")
var id = flag.String("i", "", "element id to find")

func main() {
	flag.Parse()

	if *f == "" || *id == "" {
		fmt.Fprintf(os.Stderr, "input parameter must not be null")
		flag.Usage()
	}

	outline3(*f, *id)
}

func outline3(filename, id string) error {
	data, err := os.Open(filename)
	if err != nil {
		return err
	}

	doc, err := html.Parse(data)
	if err != nil {
		return err
	}

	n := elementByID(doc, id)
	if n != nil {
		fmt.Printf("node name: %s\n", n.Data)
	}

	return nil
}

var element *html.Node
var attrid string

func elementByID(doc *html.Node, id string) *html.Node {
	attrid = id
	forEachNode(doc, startElement, endElement)

	return element
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil {
		if pre(n) {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

func startElement(n *html.Node) bool {

	if n.Type != html.ElementNode {
		return false
	}
	if n.Data != "a" {
		return false
	}

	for _, a := range n.Attr {
		if a.Key == "href" {
			element = n
			return true
		}
	}
	return false
}

func endElement(n *html.Node) bool {
	return false
}

//!-startend
