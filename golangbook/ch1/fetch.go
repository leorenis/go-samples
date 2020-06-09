// Package ch1 is Book's exemples for the Chapter One
package ch1

import (
	"fmt"
	"io"
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
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v \n", url, err)
			os.Exit(1)
		}
	}
}
