package ch3

import "fmt"

// ShowBitwise is
func ShowBitwise() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x) // 00100010, o conjunto {1. 5}
	fmt.Printf("%08b\n", y) // 00000110, o conjunto {1. 2}
}
