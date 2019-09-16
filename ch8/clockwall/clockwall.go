package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	if len(os.Args) < 1 {
		fmt.Fprintf(os.Stderr, "no time zone information")
		os.Exit(-1)
	}
	for _, t := range os.Args[1:] {
		i := strings.Index(t, "=")
		if i == -1 {
			continue
		}
		conn, err := net.Dial("tcp", t[i+1:])
		log.Printf("port:%v, %v", t[i], err)
		if err != nil {
			continue
		}
		wg.Add(1)
		go handleConn(conn, t[:i])
	}
	wg.Wait()
}

func handleConn(c net.Conn, city string) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		fmt.Printf("%s : %s\n", city, input.Text())
	}
	wg.Done()
}
