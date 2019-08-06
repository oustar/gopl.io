// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package tempconv

import "fmt"

func Example_two() {
	//!+printf
	var k Kelvin
	k = CToK(-173.15)
	fmt.Println(k.String()) // "100°C"
	fmt.Printf("%v\n", k)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", k)   // "100°C"
	fmt.Println(k)          // "100°C"
	fmt.Printf("%g\n", k)   // "100"; does not call String
	fmt.Println(float64(k)) // "100"; does not call String
	var i Kelvin
	i = -173.15 + 273.15
	fmt.Println(i)
	//!-printf

	// Output:
	// 100°C
	// 100°C
	// 100°C
	// 100°C
	// 100
	// 100
}
