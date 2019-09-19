package main

func main() {

	b := make(chan int)
	go func(ch chan int) {
		for {
			ch <- 1
		}
	}(b)

	var a chan int
	for {
		a = b
		b = make(chan int)
		go func(in, out chan int) {
			for {
				i := <-in
				out <- i
			}
		}(a, b)
	}
}
