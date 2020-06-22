package tempconv

// CToF is
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC is
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CToK is
func CToK(c Celsius) Kelvin {
	return Kelvin(c + (-273.15))
}

// KToC is
func KToC(k Kelvin) Celsius {
	return Celsius(k - (-273.15))
}

// FToK is
func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f-32)*5/9 + 273.15)
}

// KToF is
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit((k-273.15)*9/5 + 32)
}
