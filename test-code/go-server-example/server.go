package main

import (
	"fmt"
	"net/http"
)

var counter = 0

func count(w http.ResponseWriter, req *http.Request) {
	counter++
	fmt.Fprintf(w, "Counter: %d", counter)
}

func main() {
	http.HandleFunc("/", count)
	http.ListenAndServe(":8090", nil)
}
