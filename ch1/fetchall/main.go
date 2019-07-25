// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	_ "math/rand"
	"net/http"
	"os"
	_ "strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		//go fetch2(url, ch)
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
		//fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch2(url string, ch chan<- string) {
	fetch(url, ch)
	fetch(url, ch)
}
func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%v", err)
		return
	}
	/* /* f, err := os.Create(getFileName(url))
	if err != nil {
		ch <- fmt.Sprintf("creating file %s, %v", url, err)
		return
	}
	nbytes, err := io.Copy(f, resp.Body) */
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s, %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs	%7d		%s", secs, nbytes, url)
}

/* func getFileName(url string) string {
	var fileName string

	rand.Seed(time.Now().UTC().UnixNano())
	n := rand.Int()

	if strings.HasPrefix(url, "https://") {
		fileName = strings.TrimLeft(url, "https://")
	} else {
		fileName = strings.TrimLeft(url, "http://")
	}

	return fmt.Sprintf("%s%d.txt", fileName, n)
} */

//!-
