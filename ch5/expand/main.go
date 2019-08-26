package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var src = flag.String("s", "hello $foo", "source strings")
var sub = flag.String("e", "foo", "expand substring")

func main() {
	flag.Parse()

	//fmt.Fprintf(os.Stdout, "expand string: %v\n", expand(*src, *sub, transfer))
	fmt.Println(*src, *sub)
	results := expand(*src, *sub, transfer)
	n, err := fmt.Print(string(results))
	if err != nil {
		fmt.Fprintf(os.Stderr, "print %d, error : %v\n", n, err)
	} else {
		fmt.Printf("%d bytes\n", n)
	}
}

func expand(s, sub string, f func(string) string) string {
	results := ""
	replaces := f(sub)
	sub = "$" + sub
	for {
		if i := strings.Index(s, sub); i != -1 {
			results += (s[:i] + replaces)
			s = s[i+len(sub):]
		} else {
			results += s
			break
		}
	}
	return results
}

func transfer(s string) string {
	if s == "foo" {
		return "world"
	}

	return ""
}
