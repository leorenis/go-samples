package main

import (
	"fmt"
)

func main() {
	personName := "John Doe"
	appVersion := 1.1
	age := 28

	fmt.Println("The first Go lang's app written by mr.", personName, "age", age)
	fmt.Println("App Version: ", appVersion)

	fmt.Println("1- Start mentoring")
	fmt.Println("2- Show logs")
	fmt.Println("3- Exit")

	var command int
	fmt.Scan(&command)

	fmt.Println("Pointer in memory for command variable is: ", &command)
	fmt.Println("The command choosen:", command)

	switch command {
	case 1:
		fmt.Println("Mentoring")
	case 2:
		fmt.Println("Logs")
	case 0:
		fmt.Println("Exit")
	default:
		fmt.Println("Unknown command.")
	}

}
