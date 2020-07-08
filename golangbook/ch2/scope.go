package ch2

import (
	"fmt"
	"log"
	"os"
)

// ShowScope is
func ShowScope() {
	x := "hello"          // fuction scope
	for _, x := range x { // implicit for scope
		x := x + 'A' - 'a' // explicit for block scope
		fmt.Printf("%c", x)
	}
}

// ScopeInIf is
func ScopeInIf() {
	if s := anyFn(); s == 0 {
		fmt.Println(s)
	} else if y := anyOtherFn(s); s == y {
		fmt.Println(s, y)
	} else {
		fmt.Println(s, y)
	}

	fmt.Println(s, y) // Compilation error. x and y are not enabled here...
}

func anyFn() int {
	return 0
}

func anyOtherFn(s int) int {
	return s
}

// Be carefull with scope

var cwd string

func init() {
	cwd, err := os.Getwd() // NOTICE: Wrong way!
	if err != nil {
		log.Fatalf("os.GetWd failed: %v", err)
	}

	log.Printf("Working directory = %s", cwd)
}

// Tip: As cwd as err were initialized on line 44. The cwd in line 41 stay not initialized.
// Check the rigth way bellow:

var cwd2 string

func init() {
	var err error
	cwd2, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}
