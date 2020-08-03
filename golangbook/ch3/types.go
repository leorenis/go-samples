package ch3

import "fmt"

// ShowTypes is
func ShowTypes() {
	differentesTypes()
	accuracyLoss()
	octalAndHexadecimal()
	showRunas()
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

func octalAndHexadecimal() {
	o := 0777
	fmt.Printf("%d %[1]o %#[1]o \n", o) // 511 777 0777

	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X \n", x) // 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
}

func showRunas() {
	ascii := 'a'
	unicode := 'Ӕ'
	newline := '\n'

	fmt.Printf("%d %[1]c %[1]c %[1]q \n", ascii)
	fmt.Printf("%d %[1]c %[1]c %[1]q \n", unicode)
	fmt.Printf("%d %[1]q \n", newline)
}
