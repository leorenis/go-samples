package ch2

import "fmt"

// ShowFtoc is a function to show two convertions from fahrenheit to Celsius
func ShowFtoc() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g •F = %g •C \n", freezingF, fToC(freezingF)) // "32 •F = 0•C "
	fmt.Printf("%g •F = %g •C \n", boilingF, fToC(boilingF))   // "212 •F = 100•C "

	// Curiosity
	pointers()
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func pointers() {
	x := 1
	p := &x
	fmt.Println("p has a address for x: ", p, " *p pointer for x value as you can see: *p=", *p, " and x=", x)
	fmt.Printf("%d \n", p) // Var p is a pointer for a int variable

	var xs, xy int
	fmt.Println(&xs == &xs, &xs == &xy, &xs == nil)
	fmt.Println("Outside", pointerFn(), "Accessing value short declared inside:", *pointerFn())

	// Each call returns a distinct value... (a new pointer to int)
	fmt.Println("pointerFn() == pointerFn()", pointerFn() == pointerFn()) // False
}

func pointerFn() *int {
	v := 1
	fmt.Println("Inside fn:", &v)
	return &v
}
