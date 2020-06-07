// Package ch1 is Book's exemples for the Chapter One
package ch1

import (
	"bufio"
	"fmt"
	"os"
)

// ShowDup is
func ShowDup() {
	dup1()
}

// Dup1 is a function exibe o texto de toda linha que aparece mais de uma vez na entrada padrao, precedida por sua contagem.
// The scanner will stop on EOF (End Of File). Typing Ctrl-D for example will send end of file and stop the scanner.
func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		text := input.Text()
		counts[text]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
