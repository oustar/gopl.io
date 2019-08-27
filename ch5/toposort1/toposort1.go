//
package main

import (
	"fmt"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	order := topoSort(prereqs)
	for i := 0; i < len(order); i++ {
		fmt.Printf("%d:\t%s\n", i+1, order[i])
	}

}

func topoSort(m map[string][]string) map[int]string {
	var i int
	seen := make(map[string]bool)
	order := make(map[int]string)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[i] = item
				i++
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	//sort.Strings(keys)
	visitAll(keys)
	return order
}

//!-main
