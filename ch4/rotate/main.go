package main

import (
	"fmt"
)

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for i := range s {
		var t []int
		t = append(t, s...)
		rotateLeft(t, i)
		fmt.Printf("Rotate left %d ：%v\n", i, t)
	}
}

func rotateLeft(s []int, n int) {
	var buf []int
	l := len(s)
	if n > l {
		n %= l
	}
	if n == 0 {
		return
	}

	for i := range s[:n] {
		buf = append(buf, s[i])
	}

	for i, x := range s[n:] {
		s[i] = x
	}
	copy(s[l-n:], buf)
}
