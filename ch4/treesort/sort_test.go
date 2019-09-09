// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package treesort

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestSortPrint(t *testing.T) {
	var tests = []struct {
		d []int
		s string
	}{
		{[]int{4, 5, 1, 2}, "1 2 4 5"},
	}

	for _, test := range tests {
		var proot *Tree
		proot = Sort(test.d)
		if proot.String() != test.s {
			t.Errorf("Tree String() is err got %s want %s", proot.String(), test.s)
		}
	}
}
