package ch2

import "fmt"

// ShowScope is
func ShowScope() {
	x := "hello"          // fuction scope
	for _, x := range x { // implicit for scope
		x := x + 'A' - 'a' // explicit for block scope
		fmt.Printf("%c", x)
	}
}
