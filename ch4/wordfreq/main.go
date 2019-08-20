package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordcount: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}

	for k, v := range counts {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
