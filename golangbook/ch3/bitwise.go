package ch3

import "fmt"

// ShowBitwise is
func ShowBitwise() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Println(x, y)
	fmt.Printf("%08b\n", x) // 00100010, o conjunto {1. 5}
	fmt.Printf("%08b\n", y) // 00000110, o conjunto {1. 2}

	fmt.Printf("%08b\n", x&y)  // 00000010, a interseccao(1)
	fmt.Printf("%08b\n", x|y)  // 00100110, a uniao(1,2,5}
	fmt.Printf("%08b\n", x^y)  // 00100100, a diferenca simetrica(2,5}
	fmt.Printf("%08b\n", x&^y) // 00100000, a diferenca (5}. O operador &^ serve para limpeza de bits (AND NOT): Na expressao `z=x&ˆy`, cada bit de Z será 0 se o bit correspondente de Y for 1. Caso contrario, será o bit correspondente de X.
}
