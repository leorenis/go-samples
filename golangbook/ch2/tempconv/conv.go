package tempconv

// CToF is
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC is
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
