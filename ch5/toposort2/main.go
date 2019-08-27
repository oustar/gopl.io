// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136. Exercise 5.11

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"os"
	"sort"
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
	"linear algebra":        {"calculus"},
}

//!-table

//!+main
func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "toposort3: %v\n", err)
		os.Exit(1)
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// NodeState 表示node的状态
type NodeState int

// 记录状态
const (
	White NodeState = iota // 没有访问
	Gray                   // 正在访问
	Black                  // 完成访问
)

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]NodeState)
	var visitAll func(items []string) error

	visitAll = func(items []string) error {
		for _, item := range items {
			switch seen[item] {
			case White:
				seen[item] = Gray
				err := visitAll(m[item])
				if err != nil {
					return err
				}
				order = append(order, item)
				seen[item] = Black
			case Gray:
				return fmt.Errorf("exist cycling")
			case Black:
				continue
			default:
				continue
			}

		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	if err := visitAll(keys); err != nil {
		return nil, err
	}
	return order, nil
}

//!-main
