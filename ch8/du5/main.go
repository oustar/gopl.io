// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 250.

// The du3 command computes the disk usage of the files in a directory.
package main

// The du3 variant traverses all directories in parallel.
// It uses a concurrency-limiting counting semaphore
// to avoid opening too many files at once.

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

// PathSize include dir name and all files size
type PathSize struct {
	Path  string
	Sizes int64
}

// RootInfo include the number of files and the length of all files
type RootInfo struct {
	Files int64
	Bytes int64
}

//!+
func main() {
	// ...determine roots...

	//!-
	flag.Parse()

	// Determine the initial directories.
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	//!+
	// Traverse each root of the file tree in parallel.
	fileSizes := make(chan PathSize)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	//!-

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	rootInfo := make(map[string]RootInfo)
	for _, root := range roots {
		rootInfo[root] = RootInfo{0, 0}
	}
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			for root, info := range rootInfo {
				if !strings.HasPrefix(size.Path, root) {
					continue
				}
				info.Files++
				info.Bytes += size.Sizes
				rootInfo[root] = info
				break
			}
		case <-tick:
			printDiskUsage(rootInfo)
		}
	}

	printDiskUsage(rootInfo) // final totals
	//!+
	// ...select loop...
}

//!-

func printDiskUsage(rootInfo map[string]RootInfo) {
	fmt.Printf("==========================================================\n")
	for root := range rootInfo {
		fmt.Printf("%s : %d files  %.1f GB\n",
			root, rootInfo[root].Files, float64(rootInfo[root].Bytes)/1e9)
	}
	fmt.Printf("==========================================================\n")
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//!+walkDir
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- PathSize) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- PathSize{dir, entry.Size()}
		}
	}
}

//!-walkDir

//!+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
