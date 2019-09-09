// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary Tree.
package treesort

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

//Tree is the link tree
type Tree struct {
	value       int
	left, right *Tree
}

// Sort sorts values in place.
func Sort(values []int) *Tree {
	var root *Tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return root
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *Tree, value int) *Tree {
	if t == nil {
		// Equivalent to return &Tree{value: value}.
		t = new(Tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func printTree(t *Tree, w io.Writer) {
	if t != nil {
		printTree(t.left, w)
		fmt.Fprintf(w, "%d ", t.value)
		printTree(t.right, w)
	}
}
func (t *Tree) String() string {
	var buf bytes.Buffer
	printTree(t, &buf)
	return strings.TrimRight(buf.String(), " ")
}

//!-
