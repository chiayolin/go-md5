package main

import (
	"fmt"
	"net/http"
	"crypto/md5"
)

func handler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Path[1:]
	if data != "" {
		fmt.Fprintf(w, "%x", md5.Sum([]byte(data)))
	} else {
		fmt.Fprintf(w, "Usage: %s/{data}", r.Host)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
