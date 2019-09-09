package main

import (
	"fmt"
	"sort"
)

func isPalindrome(s sort.Interface) bool {
	n := s.Len()
	if n < 2 {
		return true
	}

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		} else {
			return false
		}
	}

	return true
}

func main() {
	data := []int{1, 2, 3, 2, 1}
	if b := isPalindrome(sort.IntSlice(data)); b == true {
		fmt.Println("ok")
	} else {
		fmt.Println("error")
	}

}
