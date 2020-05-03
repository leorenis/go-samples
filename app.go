package main

import (
	"fmt"
	"reflect"
)

func main() {
	personName := "John Doe"
	appVersion := 1.1
	age := 28

	fmt.Println("The first Go lang's app written by mr.", personName, "age", age)
	fmt.Println("App Version: ", appVersion)

	fmt.Println("Types: ", reflect.TypeOf(personName), reflect.TypeOf(age))
}
