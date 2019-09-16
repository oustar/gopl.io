package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	flag.Parse()
	roots := flag.Args()

	if len(roots) == 0 {
		return
	}

	var wg sync.WaitGroup
	response := make(chan string, len(roots))
	for _, r := range roots {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			resp, err := request(s)
			if err != nil {
				return
			}
			response <- resp
		}(r)
	}

	select {
	case s := <-response:
		fmt.Printf("%s\n", s)
		close(done)
	}
	wg.Wait()
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

func request(hostname string) (response string, err error) {

	req, err := http.NewRequest("GET", hostname, nil)
	if err != nil {
		return "", err
	}
	req.Cancel = done

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("getting %s: %s", hostname, resp.Status)
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	return fmt.Sprintf("%s:%d", hostname, nbytes), nil
}
