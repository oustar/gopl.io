package unitconv

import "fmt"

func Example_lenconv() {
	f := Feet(100)
	m := FtToM(f)
	fmt.Printf("%s = %s\n", f, m)

	// output:
	// 100
	// 100
}

func Example_weiconv() {
	p := Pound(100)
	k := LbToKg(p)
	fmt.Printf("%s = %s", p, k)

	// output:
	// 100
}
