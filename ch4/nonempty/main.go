// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 91.

//!+nonempty

// Nonempty is an example of an in-place slice algorithm.
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

//!-nonempty

func main() {
	//!+main
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`

	s := []string{"one", "one", "three"}
	fmt.Printf("%q\n", nodup(s))

	s1 := []byte("hello　　　世界！")
	b := noDupSpace(s1)
	fmt.Printf("len = %d, %q\n", len(b), b)
	//!-main
}

func nodup(s []string) []string {
	var i int
	for j := range s {
		if i != j {
			if s[i] != s[j] {
				i++
				if i != j {
					s[i] = s[j]
				}
			}
		}
	}

	return s[:i+1]
}

//!+alt
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func noDupSpace(u []byte) []byte {
	preSpaceFlag := false
	preSpaceIndex := -1
	i := 0
	n := 0
	count := utf8.RuneCount(u)
	for n < count {
		r, size := utf8.DecodeRune(u[i:])
		if unicode.IsSpace(r) {
			if !preSpaceFlag {
				preSpaceFlag = true
				preSpaceIndex = i
			}
			i += size
		} else {
			if preSpaceFlag {
				u[preSpaceIndex] = ' '
				copy(u[preSpaceIndex+1:], u[i:])
				i = preSpaceIndex + 1 + size

				preSpaceFlag = false
				preSpaceIndex = -1
			} else {
				i += size
			}
		}
		n++
	}
	return u[:i]
}

//!-alt
