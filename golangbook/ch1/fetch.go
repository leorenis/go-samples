// Package ch1 is Book's exemples for the Chapter One
package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Fetch is
// to run: go run book.go google.com
// out in a file: go run book.go google.com g1.globo.com youtube.com > rs.txt
func Fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(urlNormalized(url))
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
		fmt.Printf("\nHTTP status: %s", resp.Status)
	}
}

func urlNormalized(url string) string {
	prefix := "http://"
	if strings.HasPrefix(url, prefix) {
		return url
	}
	return prefix + url
}
