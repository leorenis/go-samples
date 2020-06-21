package ch2

import "fmt"

// Celsius is
type Celsius float64

// Fahrenheit is
type Fahrenheit float64

// Consts is
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	Boiling       Celsius = 100
)

// ShowTempConv is
func ShowTempConv() {
	fmt.Printf("%g \n", ctoF(10))
	fmt.Printf("%g \n", fahrenheitToC(10))
	fmt.Printf("%g \n", Boiling-FreezingC) // 100•C
	boilingF := ctoF(Boiling)
	fmt.Printf("%g \n", boilingF-ctoF(FreezingC)) // 180•C
	// fmt.Printf("%g \n", boilingF-FreezingC) // Erro de compilacao. Incompatibilidade de tipos

	// More exemples
	c := fahrenheitToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c) // Nao ha necessidade de chamar metodo string explicitamente.
	fmt.Println(c)        // Noa chama metodo String
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))

}

func ctoF(c Celsius) Fahrenheit          { return Fahrenheit(c*9/5 + 32) }
func fahrenheitToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func (c Celsius) String() string         { return fmt.Sprintf("%g °C", c) }
