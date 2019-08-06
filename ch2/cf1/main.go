// unitconv
package main

import (
	"flag"
	"fmt"
	"gopl.io/ch2/unitconv"
	"os"
	"strconv"
)

var unitptr = flag.String("u", "len", "the unit type: len(default) or wei")

func main() {
	flag.Parse()

	if *unitptr != "len" && *unitptr != "wei" {
		fmt.Fprintf(os.Stderr, "cf1: unit type error")
		os.Exit(1)
	}

	args := flag.Args()
	for _, s := range args {
		n, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf1: input para err")
			os.Exit(1)
		}

		if *unitptr == "len" {
			f := unitconv.Feet(n)
			fmt.Fprintf(os.Stdout, "%s = %s\n", f, unitconv.FtToM(f))
		} else if *unitptr == "wei" {
			p := unitconv.Pound(n)
			fmt.Fprintf(os.Stdout, "%s = %s\n", p, unitconv.LbToKg(p))
		}
	}
}
