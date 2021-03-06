package ch3

import "fmt"

// ShowBitwise is
func ShowBitwise() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Println(x, y)
	fmt.Printf("%08b\n", x) // 00100010, o conjunto {1. 5} -> Leia elevado, ex: {2 elevado 1, 2 elevado 5}.
	fmt.Printf("%08b\n", y) // 00000110, o conjunto {1. 2}

	fmt.Printf("%08b\n", x&y)  // 00000010, a interseccao(1)
	fmt.Printf("%08b\n", x|y)  // 00100110, a uniao(1,2,5}
	fmt.Printf("%08b\n", x^y)  // 00100100, a diferenca simetrica(2,5}
	fmt.Printf("%08b\n", x&^y) // 00100000, a diferenca (5}. O operador &^ serve para limpeza de bits (AND NOT): Na expressao `z=x&ˆy`, cada bit de Z será 0 se o bit correspondente de Y for 1. Caso contrario, será o bit correspondente de X.

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i) // `1`, `5`
		}
	}

	fmt.Printf("%08b\n", x<<1) // 01000100, o conjunto {2, 6}
	fmt.Printf("%08b\n", x>>1) // 00010001, o conjunto {0, 4}

}

// Code same page of book, for this reason this code is here.
func medals() {
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) // "bronze", "silver", "gold"
	}
}

// Note: len() devolve um int com sinal.
//       Se len devolvesse um numero sem sinal e.g. uint. Seria algo desastroso,
//       Sendo que na terceira iteracao, em que i == 0, a instrucao i-- faria i se tornar
//       o valor maximo de uint (por exemplo: (2 elevado 64) -1)
