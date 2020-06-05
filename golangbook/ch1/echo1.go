// Book's exemples for the Chapter One
package ch1

import (
	"fmt"
	"os"
)

// SliceRagesTesteds is
func SliceRagesTesteds() {
	s := []string{"Abobora", "tomate", "manga"}
	fmt.Println(s[0:3])
}

// Echo1 is
func Echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// Echo2 is
func Echo2() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
