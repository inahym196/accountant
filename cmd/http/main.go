package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	fmt.Fprintf(w, "%s", body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8080", nil)
}
