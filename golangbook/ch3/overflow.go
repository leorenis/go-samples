package ch3

import "fmt"

// ShowOverflow is
func ShowOverflow() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
}
