package ch2

import "fmt"

// ShowVariables is
func ShowVariables() {
	// Curiosity
	pointers()
}

func pointers() {
	x := 1
	p := &x
	fmt.Println("p has a address for x: ", p, " *p pointer for x value as you can see: *p=", *p, " and x=", x)
	fmt.Printf("%d \n", p) // Var p is a pointer for a int variable
	fmt.Println("Samething still crazy: ", &*p, &*p == p)

	var xs, xy int
	fmt.Println(&xs == &xs, &xs == &xy, &xs == nil)
	fmt.Println("Outside", pointerFn(), "Accessing value short declared inside:", *pointerFn())

	// Each call returns a distinct value... (a new pointer to int)
	fmt.Println("pointerFn() == pointerFn()", pointerFn() == pointerFn()) // False

	v := 1
	inc(&v)              // Side effect: v now is 2
	fmt.Println(inc(&v)) // "3" (and v is 3)
}

func pointerFn() *int {
	v := 1
	fmt.Println("Inside fn:", &v)
	return &v
}

func inc(p *int) int {
	*p++
	return *p
}
