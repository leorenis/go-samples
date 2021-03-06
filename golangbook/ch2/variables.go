package ch2

import "fmt"

// ShowVariables is
func ShowVariables() {
	// Curiosity
	pointers()
	newFn()

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

func newFn() {
	p := new(int) // p, *int type, pointers to annonymous variable int.
	fmt.Println("The int without name with 0 `Initial value to int`", *p)

	*p = 2
	fmt.Println("The int without name with:", *p)

	// Each call to `new` returns a diferent variable with a different address.
	a := new(int) // create a new int type variable without name.
	b := new(int)
	fmt.Println(a == b) // false
}

// The two functions bellow are the same... only written of diferents way.
func fnNew() *int {
	return new(int)
}

func funcfn2New() *int {
	var dummy int
	return &dummy
}
