package main

import (
	"fmt"
	"testing"
)

func ExampleExpand() {
	const src = "hello $foo"
	const sub = "foo"
	fmt.Printf("src:%s,expand:%s\n", src, expand(src, sub, transfer))

	// Output:
	// src:hello $foo,expand:hello world
}

func TestExpand(t *testing.T) {
	var tests = []struct {
		src  string
		sub  string
		want string
	}{
		{"hello $foo", "foo", "hello world"},
		{"$foo hello", "foo", "world hello"},
		{"hello $foo OK", "foo", "hello world OK"},
	}
	for _, test := range tests {
		got := expand(test.src, test.sub, transfer)
		if got != test.want {
			t.Errorf("%q = expand(%q, %q), want %q", got, test.src, test.sub, test.want)
		}
	}
}
