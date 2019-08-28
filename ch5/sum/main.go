// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import "fmt"

//!+
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(n int, vals ...int) int {
	nmax := n
	for _, val := range vals {
		if val > nmax {
			nmax = val
		}
	}
	return nmax
}

func min(n int, vals ...int) int {
	nmin := n
	for _, val := range vals {
		if val < nmin {
			nmin = val
		}
	}
	return nmin
}

//!-

func main() {
	//!+main
	fmt.Println(sum())           //  "0"
	fmt.Println(sum(3))          //  "3"
	fmt.Println(sum(1, 2, 3, 4)) //  "10"
	//!-main

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
	//!-slice

	fmt.Println(max(1))
	fmt.Println(max(1, values...))

}
