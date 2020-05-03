package main

import (
	"fmt"
	"os"
)

func main() {

	showIntroduction()
	showMenu()

	command := readCommand()
	switch command {
	case 1:
		fmt.Println("Mentoring")
	case 2:
		fmt.Println("Logs")
	case 0:
		os.Exit(0)
	default:
		os.Exit(-1)
	}

}

func showIntroduction() {
	personName := "John Doe"
	appVersion := 1.1
	age := 28

	fmt.Println("The first Go lang's app written by mr.", personName, "age", age)
	fmt.Println("App Version: ", appVersion)
}

func showMenu() {
	fmt.Println("1- Start mentoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The command choosen:", command)
	return command
}
