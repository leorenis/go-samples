// Package ch1 is Book's exemples for the Chapter One
package ch1

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// ShowDup is
func ShowDup() {
	dup1()
}

// Dup1 is a function exibe o texto de toda linha que aparece mais de uma vez na entrada padrao, precedida por sua contagem.
func dup1() {
	start := time.Now()
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	fmt.Printf("%dns eleapsed \n", time.Since(start).Nanoseconds())
}
