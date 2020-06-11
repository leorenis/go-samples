// Package ch1 is Book's exemples for the Chapter One
package ch1

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// FetchAll is
// to run: go run book.go google.com g1.globo.com youtube.com
func FetchAll() {
	start := time.Now()

	if len(os.Args) <= 1 {
		portals := portals()
		process(portals)
	} else {
		process(os.Args[1:])
	}

	fmt.Printf("%.2fs eleapsed \n", time.Since(start).Seconds())
}

func process(urls []string) {
	ch := make(chan string)

	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch) // chanel receive
	}
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

func portals() []string {
	file, err := os.Open("ch1/resources/top-domains.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error to read file: %v\n", err)
	}

	reader := bufio.NewReader(file)
	var portals []string
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		portals = append(portals, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return portals
}
