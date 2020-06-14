package ch2

import "fmt"

// ShowFtoc is a function to show two convertions from fahrenheit to Celsius
func ShowFtoc() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g •F = %g •C \n", freezingF, fToC(freezingF)) // "32 •F = 0•C "
	fmt.Printf("%g •F = %g •C \n", boilingF, fToC(boilingF))   // "212 •F = 100•C "
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
