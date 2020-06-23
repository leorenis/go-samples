package generalconv

import "fmt"

// Feet is
type Feet float64

// Metre is
type Metre float64

func (feet Feet) String() string {
	return fmt.Sprintf("%g Feet", feet)
}

func (m Metre) String() string {
	return fmt.Sprintf("%g Meters", m)
}
