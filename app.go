package main

import (
	"fmt"
	"reflect"
)

func main() {
	var personName = "LEO R R Santos"
	var appVersion float32 = 1.1
	var age = 28

	fmt.Println("The first Go lang's app written by mr.", personName, "age", age)
	fmt.Println("App Version: ", appVersion)

	fmt.Println("Types: ", reflect.TypeOf(personName), reflect.TypeOf(age))
}
