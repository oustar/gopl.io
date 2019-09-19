package main

import (
	"fmt"
	"time"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	a := make(chan int64)
	b := make(chan int64)

	var n int64

	go func() {
		for {
			if cancelled() {
				close(b)
				break
			}
			i := <-a
			i++
			b <- i
		}
	}()

	go func() {
		for {
			if cancelled() {
				close(a)
				break
			}
			j := <-b
			j++
			n = j
			a <- j

		}
	}()

	a <- 0
	time.Sleep(1 * time.Second)
	close(done)

	fmt.Printf("%d times", n)

}
