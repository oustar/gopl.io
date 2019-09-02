package main

import (
	"fmt"

	"gopl.io/ch6/geometry"
)

func main() {
	perim := geometry.Path{{X: 1, Y: 1},
		{X: 5, Y: 1}, {X: 5, Y: 4}, {X: 1, Y: 1}}
	fmt.Println(perim.Distance())
}
