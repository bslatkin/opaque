package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyHandler struct {}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Easiest webserver ever")
}

func main() {
	http.Handle("/", &MyHandler{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
