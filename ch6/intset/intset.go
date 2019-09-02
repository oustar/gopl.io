// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Len calculate the lenght of IntSet
func (s *IntSet) Len() int {
	popCount := func(x uint64) int {
		n := 0
		for x != 0 {
			x = x & (x - 1) // clear rightmost non-zero bit
			n++
		}
		return n
	}
	l := 0
	for _, word := range s.words {
		l += popCount(word)
	}
	return l
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds the slices of int
func (s *IntSet) AddAll(xs []int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// Remove remove x form the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &^= 1 << bit
	}
}

// Clear remove all elements form the set
func (s *IntSet) Clear() {
	s.words = []uint64{}
}

// Copy return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var t IntSet
	t.words = append(t.words, s.words...)
	return &t
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith set s to the intersect of s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	l := len(s.words)
	if l > len(t.words) {
		l = len(t.words)
	}
	for i := 0; i < l; i++ {
		s.words[i] &= t.words[i]
	}
	s.words = s.words[0:l]
}

// DifferenceWith set s to the difference of s and t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, sword := range s.words {
		if i >= len(t.words) {
			break
		}
		for j := 0; j < 64; j++ {
			if (sword&(1<<uint(j)) != 0) && (t.words[i]&(1<<uint(j)) != 0) {
				s.words[i] &^= (1 << uint(j))
			}
		}
	}
}

// SymmetricDifference set s to the symmetric difference of s and t
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Elems return the of slices of IntSet
func (s *IntSet) Elems() []int {
	var a []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				a = append(a, 64*i+j)
			}
		}
	}
	return a
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string
