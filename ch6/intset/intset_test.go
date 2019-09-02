// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestLen(t *testing.T) {
	var tests = []struct {
		s    []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, 2, 65}, 3},
		{[]int{1, 2, 65, 129}, 4},
	}
	for _, test := range tests {
		var set IntSet
		for _, n := range test.s {
			set.Add(n)
		}
		if set.Len() != test.want {
			t.Errorf("IntSet.len() is err, got %d, want %d", set.Len(), test.want)
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		s   []int
		r   int
		str string
	}{
		{[]int{}, 1, "{}"},
		{[]int{0, 1, 2}, 1, "{0 2}"},
		{[]int{0, 65, 128}, 128, "{0 65}"},
	}

	for _, test := range tests {
		var set IntSet
		for _, i := range test.s {
			set.Add(i)
		}
		set.Remove(test.r)

		if set.String() != test.str {
			t.Errorf("IntSet.Remove is err, got %s want %s", set.String(), test.str)
		}
	}
}

func TestClear(t *testing.T) {
	var tests = []struct {
		s   []int
		str string
	}{
		{[]int{1}, "{}"},
		{[]int{1, 67}, "{}"},
		{[]int{1, 67, 129}, "{}"},
		{[]int{}, "{}"},
	}

	for _, test := range tests {
		var set IntSet
		set.AddAll(test.s)
		set.Clear()
		if set.String() != test.str {
			t.Errorf("IntSet.Clear is err, got %s, want %s", set.String(), test.str)
		}
	}
}

func TestCopy(t *testing.T) {
	var tests = []struct {
		s   []int
		str string
	}{
		{[]int{}, "{}"},
		{[]int{1}, "{1}"},
		{[]int{1, 67}, "{1 67}"},
		{[]int{1, 67, 129}, "{1 67 129}"},
	}

	for _, test := range tests {
		var a IntSet
		a.AddAll(test.s)
		pb := a.Copy()

		if (*pb).String() != test.str {
			t.Errorf("IntSet.Copy is err, got %s want %s", (*pb).String(), test.str)
		}
	}
}

func TestIntersectWith(t *testing.T) {
	var tests = []struct {
		s, t []int
		str  string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{1}, []int{1}, "{1}"},
		{[]int{1, 63, 129}, []int{1, 64, 129}, "{1 129}"},
	}
	for _, test := range tests {
		var a, b IntSet
		a.AddAll(test.s)
		b.AddAll(test.t)
		a.IntersectWith(&b)
		if a.String() != test.str {
			t.Errorf("IntSet.IntersectWith is err got %s want %s", a.String(), test.str)
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	var tests = []struct {
		s, t []int
		str  string
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, "{}"},
		{[]int{1, 2, 3}, []int{2}, "{1 3}"},
		{[]int{1, 65, 129, 512}, []int{1, 2, 1025}, "{65 129 512}"},
	}

	for _, test := range tests {
		var a, b IntSet
		a.AddAll(test.s)
		b.AddAll(test.t)
		a.DifferenceWith(&b)
		if a.String() != test.str {
			t.Errorf("IntSet.Difference is err, got %s want %s", a.String(), test.str)
		}
	}
}

func TestSymmetricDifferenc(t *testing.T) {
	var tests = []struct {
		s, t []int
		str  string
	}{
		{[]int{1, 2}, []int{5, 6}, "{1 2 5 6}"},
		{[]int{1, 2, 65, 129}, []int{5, 6}, "{1 2 5 6 65 129}"},
		{[]int{1, 2}, []int{5, 6, 65, 129, 513}, "{1 2 5 6 65 129 513}"},
		{[]int{1, 2, 129, 1025}, []int{5, 6, 65, 129, 513}, "{1 2 5 6 65 513 1025}"},
	}

	for _, test := range tests {
		var a, b IntSet
		a.AddAll(test.s)
		b.AddAll(test.t)
		a.SymmetricDifference(&b)
		if a.String() != test.str {
			t.Errorf("IntSet.SymmetricDifference is err, got %s want %s", a.String(), test.str)
		}
	}
}

func TestElems(t *testing.T) {
	var tests = []struct {
		s, t []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 65, 129}, []int{1, 65, 129}},
		{[]int{1, 1025}, []int{1, 1025}},
	}
	var compare = func(s, t []int) bool {
		if len(s) != len(t) {
			return false
		}
		for i := range s {
			if s[i] != t[i] {
				return false
			}
		}
		return true
	}

	for _, test := range tests {
		var a IntSet
		a.AddAll(test.s)

		if !compare(a.Elems(), test.t) {
			t.Errorf("IntSet.Elems is err, got %q, want %q", a.Elems(), test.t)
		}
	}
}
