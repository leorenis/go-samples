// Package ch2 is used to write all programs showed in chapter 2 of the The Go Programming Language's book.
package ch2

import "fmt"

const boilingF = 212.0

// ShowBoiling is
func ShowBoiling() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Boiling point = %g F or %g C\n", f, c)
}
