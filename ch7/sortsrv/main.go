package main

import (
	"log"
	"net/http"
	"os"
	"sort"
)

var l *log.Logger

func init() {
	l = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
}
func main() {
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sortflag := r.URL.Query().Get("sort")
	l.Printf("?Sort=%s\n", sortflag)
	switch sortflag {
	case "Artist":
		sort.Sort(byArtist(tracks))
	case "Year":
		sort.Sort(byYear(tracks))
	}
	htmlPrint(w, tracks)
}
