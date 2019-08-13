package main

import (
	"fmt"
)

func Example_commaInt() {
	fmt.Println(commaInt("123456")) //123,456
	fmt.Println(commaInt("1234"))
	fmt.Println(commaInt("1234567"))

	// Output:
	// 123,456
	// 1,234
	// 1,234,567
}

func Example_commaFloat() {
	fmt.Println(commaFloat("-123456.78")) //123,456
	fmt.Println(commaFloat("1234"))
	fmt.Println(commaFloat("-12.34567"))

	// Output:
	// -123,456.78
	// 1,234
	// -12.34567
}

func Example_isAnagrams() {
	fmt.Println(isAnagrams("hello", "lloeh")) // true
	fmt.Println(isAnagrams("hello", "lloh"))
	// Output:
	// true
	// false
}
