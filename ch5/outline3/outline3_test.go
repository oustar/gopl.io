package main

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestOutline3(t *testing.T) {
	const filename = "./gopl.html"
	data, err := os.Open(filename)
	if err != nil {
		t.Errorf("read testing file error (%v)", err)
	}

	doc, err := html.Parse(data)
	if err != nil {
		t.Errorf("parse html file error (%v)", err)
	}

	//!+call
	forEachNode(doc, startElement, endElement)

}

func Exampleoutline() {
	const filename = "./test1.html"
	data, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read testing file error (%v)", err)
	}

	doc, err := html.Parse(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse html file error (%v)", err)
	}

	//!+call
	forEachNode(doc, startElement, endElement)

	// Output:
	// true
}
