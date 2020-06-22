// Package tempconv converts temperatures in severals scales.
package tempconv

import "fmt"

// Celsius is
type Celsius float64

// Fahrenheit is
type Fahrenheit float64

// Kelvin is
type Kelvin float64

// Consts is
const (
	AbsoluteZeroC Celsius = -273.15
	ZeroK         Kelvin  = -273.15
	FreezingC     Celsius = 0
	Boiling       Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g °C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g °F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g °K", k)
}
