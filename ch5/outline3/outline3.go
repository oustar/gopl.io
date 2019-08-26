// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 134. Exercise 5.7

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline3(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

func outline3(filename string) error {
	data, err := os.Open(filename)
	if err != nil {
		return err
	}

	doc, err := html.Parse(data)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node) {

	if n.Type == html.TextNode && n.Data != "" {
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
		depth++
	}
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s ", depth*2, "", n.Data)
		if n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Printf("href='%s' ", a.Val)
				}
			}
		}
		fmt.Printf(">\n")
		depth++
	}

}

func endElement(n *html.Node) {
	if n.Type == html.TextNode && n.Data != "" {
		depth--
	}
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

//!-startend
