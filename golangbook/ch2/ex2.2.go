package ch2

import (
	"bufio"
	"fmt"
	"gosamples/golangbook/ch2/tempconv"
	"os"
	"strconv"
)

// ShowEx2Dot4 is a function to call general convertions.
func ShowEx2Dot4() {

	processInput()
}

func processInput() {
	if len(os.Args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			text := input.Text()
			converter(text)
		}
	} else {
		for _, arg := range os.Args[1:] {
			converter(arg)
		}
	}
}

func converter(arg string) {
	number, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}

	// Prints
	fmt.Println(tempconv.FToC(tempconv.Fahrenheit(number)))
}
