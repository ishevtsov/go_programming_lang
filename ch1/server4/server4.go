package main

import (
	"ch1/server4/lissajous"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous.Lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
