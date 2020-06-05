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
