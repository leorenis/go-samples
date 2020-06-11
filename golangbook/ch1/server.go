package ch1

import (
	"fmt"
	"log"
	"net/http"
)

// StartServer is
func StartServer() {
	server1()
}

func server1() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
