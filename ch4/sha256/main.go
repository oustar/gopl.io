// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
)

//!+
import (
	"crypto/sha256"
	"crypto/sha512"
)

var shaLen = flag.Int("l", 256, "the length of SHA2")

func main() {
	flag.Parse()
	if *shaLen != 256 && *shaLen != 384 && *shaLen != 512 {
		*shaLen = 256
	}

	input := bufio.NewScanner(os.Stdin)
	var s string
	for input.Scan() {
		s += input.Text()
	}

	fmt.Printf("\nSrc:%s\n", s)
	switch *shaLen {
	case 256:
		fmt.Printf("SHA256:%X\n", sha256.Sum256([]byte(s)))
	case 384:
		fmt.Printf("SHA384:%X\n", sha512.Sum384([]byte(s)))
	case 512:
		fmt.Printf("SHA512:%X\n", sha512.Sum512([]byte(s)))
	default:
		fmt.Printf("Input no data!\n")
	}

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("The difference is %f bits\n", math.Abs(float64(PopCountSha256(c1)-PopCountSha256(c2))))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

// PopCountSha256 is the function
func PopCountSha256(c [32]byte) int {
	var n int
	for i := range c {
		n += PopCountByClearing(c[i])
	}
	return n
}

// PopCountByClearing is the function
func PopCountByClearing(x byte) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

//!-
