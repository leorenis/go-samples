// Package ch1 is Book's exemples for the Chapter One
package ch1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// CustomFetchAll is
// to run: go run book.go google.com g1.globo.com youtube.com
func CustomFetchAll() {
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
