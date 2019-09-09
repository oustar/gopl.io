// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 214.
//!+

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	dec := xml.NewDecoder(bytes.NewReader(content))

	root, _ := Parse(dec)

	Print(os.Stdout, root)
}

//!-
