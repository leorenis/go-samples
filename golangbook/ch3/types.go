package ch3

import "fmt"

// ShowTypes is
func ShowTypes() {
	differentesTypes()
	accuracyLoss()
}

func differentesTypes() {
	var apples int32 = 1
	var oranges int16 = 2
	// var compote int = apples + oranges <- Invalid operation (mismatched types int32 and int16)

	// To solve: Use conversor
	var compote = int(apples) + int(oranges)
	fmt.Println(compote)
}

func accuracyLoss() {
	f := 3.141
	i := int(f)
	fmt.Println(f, i)
	f = 1.99
	fmt.Println(int(f))

	f = 1e100
	i = int(f)
	fmt.Println(f, i)
}

// Note:
//      Conversoes de ponto flutuante para inteiro descartam qualquer parte fracionaria.
//      Truncando o resultado na direcao do zero.
