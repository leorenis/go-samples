package ch3

import "fmt"

// ShowTypes is
func ShowTypes() {
	differentesTypes()
}

func differentesTypes() {
	var apples int32 = 1
	var oranges int16 = 2
	// var compote int = apples + oranges <- Invalid operation (mismatched types int32 and int16)

	// To solve: Use conversor
	var compote = int(apples) + int(oranges)
	fmt.Println(compote)
}
