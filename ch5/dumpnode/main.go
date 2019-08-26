package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

//!+
func main() {
	for _, url := range os.Args[1:] {
		err := dumpURL(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dumpnode: %v\n", err)
			continue
		}
	}
}

func dumpURL(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	dumpNode(doc)
	return nil
}

func dumpNode(n *html.Node) {
	fmt.Printf("Node type: %v\tNode data:%s\n", n.Type, n.Data)
	for _, a := range n.Attr {
		fmt.Printf("\t%s\t%s\n", a.Key, a.Val)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dumpNode(c)
	}
}
