// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139. Exercise 5.13

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"gopl.io/ch5/links"
)

//!+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

// 主要是为了下载FIDO规范进行了改写，部分代码不具有通用性
func crawl(u string) []string {
	fmt.Println(u)
	list, err := links.Extract(u)
	if err != nil {
		log.Print(err)
	}
	var results []string
	for _, s := range list {
		if !strings.HasPrefix(s, u) {
			continue
		}
		rel, err := url.Parse(s)
		if err != nil {
			continue
		}
		if rel.RawQuery != "" {
			continue
		}
		dir, file := filepath.Split(rel.Path)
		if file != "" {
			getFile(s, dir, file)
			continue
		}
		results = append(results, s)
	}
	return results
}

func getFile(s, dir, file string) error {
	resp, err := http.Get(s)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getting %s: %s", s, resp.Status)
	}

	if err := os.MkdirAll("./"+dir, os.ModePerm); err != nil {
		return fmt.Errorf("create dir error (%v)", err)
	}

	f, err := os.OpenFile("./"+dir+file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := bufio.NewReaderSize(resp.Body, 1024*32)
	writer := bufio.NewWriter(f)

	buff := make([]byte, 32*1024)
	written := 0

	for {
		nr, er := reader.Read(buff)
		if nr > 0 {
			nw, ew := writer.Write(buff[0:nr])
			if nw > 0 {
				written += nw
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}

	return err
}

//!-crawl

//!+main
func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "command line para error")
		os.Exit(1)
	}

	breadthFirst(crawl, os.Args[1:])
	fmt.Println("OK")
	// for _, arg := range os.Args[1:] {
	// 	for _, url := range crawl(arg) {
	// 		fmt.Println(url)
	// 	}
	// }
}

//!-main
