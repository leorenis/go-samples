package ch1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

// StartServer is
func StartServer() {
	// server1()
	// server3()
	serverLissajous()
}

func server1() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", handlerCount)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func handlerCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d \n", count)
	mu.Unlock()
}

func server3() {
	http.HandleFunc("/", handlerServer3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handlerServer3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q \n", k, v)
	}
	fmt.Fprintf(w, "Host = %q \n", r.Host)
	fmt.Fprintf(w, "Remote Addr = %q \n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q = %q\n", k, v)
	}
}

func serverLissajous() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles := r.URL.Query().Get("cycles")
		nCycles, err := strconv.ParseFloat(cycles, 64)
		if err != nil {
			lissajous(w, 20)
			log.Printf("%s", err)
		} else {
			lissajous(w, nCycles)
		}

	})
	http.ListenAndServe("localhost:8000", nil)
}
