package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var handler http.HandlerFunc
	handler = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Easiest webserver ever")
	}
	http.Handle("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
