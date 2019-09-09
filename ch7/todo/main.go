package main

import (
	"log"
	"net/http"
)

func main() {
	mux := NewMux()
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
