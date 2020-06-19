package ch2

// Celsius is
type Celsius float64

// Fahrenheit is
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	Boiling       Celsius = 100
)

// ShowTempConv is
func ShowTempConv() {
	ctoF(10)
	fahrenheitToC(10)
}

func ctoF(c Celsius) Fahrenheit          { return Fahrenheit(c*9/5 + 32) }
func fahrenheitToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
