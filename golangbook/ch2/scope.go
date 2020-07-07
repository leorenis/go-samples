package ch2

import "fmt"

// ShowScope is
func ShowScope() {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x)
	}
}
