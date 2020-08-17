package ch3

import (
	"fmt"
	"math"
)

// ShowTypes is
func ShowTypes() {
	differentesTypes()
	accuracyLoss()
	octalAndHexadecimal()
	showRunas()

	showFloats()
	doubtfulDivisors()
	showNaN()
	compute()

	showComplexTypes()
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

func showFloats() {
	var f float32 = 1 << 24 // 16777216
	fmt.Printf("%f\n", f)
	fmt.Println(f) // notacao cientifica: https://bit.ly/39YABAw

	fmt.Printf("%f == %f = ", f, f+1)
	fmt.Println(f == f+1)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}

func doubtfulDivisors() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
}

func showNaN() {
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)
}

// Note:
// If a function returns a float number fails, the best thing to do is inform about error separately, like exemple bellow.

func compute() (value float64, ok bool) {
	failed := true
	//...
	if failed {
		return 0, false
	}
	result := .999
	return result, true
}

func showComplexTypes() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
}
