// Package ch1 is Book's exemples for the Chapter One
package ch1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// FetchAll is
// to run: go run book.go google.com g1.globo.com youtube.com
func FetchAll() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // chanel receive
	}

	fmt.Printf("%.2fs eleapsed \n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(urlNormalized(url))
	if err != nil {
		ch <- fmt.Sprint(err) // Send to chanel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // avoid leak resource
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v\n", url, err)
		return
	}

	msecs := time.Since(start).Milliseconds()
	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%d \t %.2fs \t %7d \t %s", msecs, secs, nbytes, url)
}
