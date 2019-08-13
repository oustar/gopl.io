// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaInt(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	// 处理头部数据
	i := n % 3
	if i == 0 {
		i = 3
	}

	var buf bytes.Buffer
	buf.WriteString(s[:i])
	for ; i < n; i += 3 {
		buf.WriteString("," + s[i:i+3])
	}
	return buf.String()
}

func commaFloat(s string) string {
	var ipre, isuf int
	if s[0] == '+' || s[0] == '-' {
		ipre = 1
	}

	isuf = strings.LastIndex(s, ".")
	if isuf == -1 {
		isuf = len(s)
	}

	return s[:ipre] + commaInt(s[ipre:isuf]) + s[isuf:]
}

func isAnagrams(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	acounts := make(map[rune]int)
	bcounts := make(map[rune]int)

	for _, r := range a {
		acounts[r]++
	}
	for _, r := range b {
		bcounts[r]++
	}

	if len(acounts) != len(bcounts) {
		return false
	}

	for r, n := range acounts {
		if bcounts[r] != n {
			return false
		}
	}
	return true
}

//!-
