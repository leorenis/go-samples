package ch3

import "fmt"

// ShowOverflow is
func ShowOverflow() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u) // Out: 255 0 1

	var i int8 = 127
	fmt.Println(i, i+1, i*i) // 127 -128 1

}
