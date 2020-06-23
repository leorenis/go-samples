package generalconv

import "fmt"

// Feet is
type Feet float64

// Metre is
type Metre float64

// Libra is
type Libra float64

// Kilograma is
type Kilograma float64

func (feet Feet) String() string {
	return fmt.Sprintf("%g Feet", feet)
}

func (m Metre) String() string {
	return fmt.Sprintf("%g Meters", m)
}

func (kg Kilograma) String() string {
	return fmt.Sprintf("%g Kg", kg)
}
