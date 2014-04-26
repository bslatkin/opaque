package main

import (
	"fmt"
	"log"
	"net/http"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Easiest webserver ever")
}

func main() {
	http.Handle("/", http.HandlerFunc(MyHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
