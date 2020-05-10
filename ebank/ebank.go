package main

import "fmt"

type CurrentAccount struct {
	holder  string
	agency  int
	number  int
	balance float64
}

func main() {
	petersAccount := CurrentAccount{"John Doe", 9889, 73784, 199909.909}
	leosAccount := CurrentAccount{"Leo Doe", 343, 3434, 10000000.909}
	fmt.Println(petersAccount)
	fmt.Println(leosAccount)

	// Other way to create a object
	var marysAccount *CurrentAccount
	var jonysAccount *CurrentAccount
	marysAccount = new(CurrentAccount)
	jonysAccount = new(CurrentAccount)
	marysAccount.holder = "Mary Doe"
	jonysAccount.holder = "Jony Bale"

	fmt.Println(*jonysAccount)
	fmt.Println(marysAccount)
}
