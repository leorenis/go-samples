package ch2

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "Separator")

// ShowEchoFourth is
// go run book.go -n -s -  Hi gente  OR go run book.go -h
func ShowEchoFourth() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
