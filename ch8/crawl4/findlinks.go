// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// WorkList include list and depth
type WorkList struct {
	list  []string
	depth int
}

// UnseenLinks include link and depth
type UnseenLinks struct {
	link  string
	depth int
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

var depth = flag.Int("depth", 3, "the depth of finding links")

//!+
func main() {

	flag.Parse()
	worklist := make(chan WorkList)       // lists of URLs, may have duplicates
	unseenLinks := make(chan UnseenLinks) // de-duplicated URLs
	var n int
	// cancel
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()
	// Add command-line arguments to worklist.
	n++
	go func() { worklist <- WorkList{flag.Args(), 0} }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				if cancelled() {
					return
				}
				var wl WorkList

				wl.list = crawl(link.link)
				wl.depth = link.depth + 1
				go func() { worklist <- wl }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		if cancelled() {
			break
		}
		list := <-worklist
		for _, link := range list.list {
			if !seen[link] {

				seen[link] = true
				if list.depth < *depth {
					n++
					unseenLinks <- UnseenLinks{link, list.depth}
				}

			}
		}

	}
}

//!-
