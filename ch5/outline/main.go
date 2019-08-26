// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//!+
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	// outline(nil, doc)
	// fmt.Printf("Name\tcounts\n")
	// for k, v := range count(doc) {
	// 	fmt.Printf("%s\t%d\n", k, v)
	// }
	printText(doc)

}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func count(n *html.Node) map[string]int {
	counts := make(map[string]int)
	if n.Type == html.ElementNode {
		counts[n.Data]++

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ctemp := count(c)
		for k, v := range ctemp {
			counts[k] += v
		}
	}
	return counts
}

func printText(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Printf("Node data: %v\n", n.Data)
		for _, v := range n.Attr {
			fmt.Printf("\t%s\t%s\n", v.Key, v.Val)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printText(c)
	}
}

//!-
