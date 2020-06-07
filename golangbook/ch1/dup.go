// Package ch1 is Book's exemples for the Chapter One
package ch1

import (
	"bufio"
	"fmt"
	"os"
)

// ShowDup is
func ShowDup() {
	dup2()
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

// Dup2 is
// Exibe a contagem e o texto das linhas que aparecem mais de uma vez na entrada. Ele lê StdIn ou uma lista de arquivos nomeados.
// Exercise 1.4 | Page. 37
func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		fileAndTextLineComputed := f.Name() + " " + input.Text()
		counts[fileAndTextLineComputed]++
	}
}
