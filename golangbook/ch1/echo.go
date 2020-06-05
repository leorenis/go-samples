// Package ch1 is Book's exemples for the Chapter One
package ch1

import (
	"fmt"
	"os"
	"strings"
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
		s += sep + os.Args[i] // ineficient way. Bad
		sep = " "
	}
	fmt.Println(s)
}

// Echo2 is
func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg // ineficient way. Bad
		sep = " "
	}
	fmt.Println(s)
}

// Echo3 is
func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " ")) // Eficient Way :) GOOD
	fmt.Println(os.Args[1:])                    // Println way
}
