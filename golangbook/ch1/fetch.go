// Package ch1 is Book's exemples for the Chapter One
package ch1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Fetch is
func Fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v \n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
