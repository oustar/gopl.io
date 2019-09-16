// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	var ok bool

	retry := make(chan struct{})

	go func() {
		for input.Scan() {
			wg.Add(1)
			go func() {
				echo(c, input.Text(), 1*time.Second)
				wg.Done()
			}()
			retry <- struct{}{}
		}
	}()

	ok = true
	for {

		select {
		case <-time.After((10 * time.Second)):
			ok = false
			log.Println("退出")

		case <-retry:
			log.Println("重新计时")

		}

		if !ok {
			break
		}

	}

	wg.Wait()
	// NOTE: ignoring potential errors from input.Err()

	c.Close()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
