package ch2

import (
	"fmt"
	"strconv"
	"strings"
)

var numbers []string // package scope level

// ShowTuples is
func ShowTuples() {
	fmt.Println(gcd(20, 32))
	fmt.Println("X:", fib(15))
	fmt.Println(strings.Join(numbers, " "))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y // Tuple assignment
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		numbers = append(numbers, strconv.Itoa(x))
	}
	return x
}
