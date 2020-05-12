package main

import "fmt"

// CurrentAccount is
type CurrentAccount struct {
	holder  string
	agency  int
	number  int
	balance float64
}

func main() {
	// testsPointers()
	mockAccount := CurrentAccount{"John Mock Doe", 343, 3434, 500.}
	mock2Account := CurrentAccount{"Mock Ac Two", 343, 3434, 1200}
	fmt.Println("Withdraw:", mockAccount.balance, mock2Account.balance)
	mockAccount.Withdraw(200)
	mock2Account.Withdraw(200)
	fmt.Println("Withdraw:", mock2Account.balance, mock2Account.balance)
}

// ( c *CurrentAccount is to similar java `this` )
func (c *CurrentAccount) Withdraw(amount float64) float64 {
	canWithDraw := amount >= 0 && amount <= c.balance
	if canWithDraw {
		c.balance -= amount
	}
	return c.balance
}

func testsPointers() {
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

	fmt.Println(*jonysAccount, jonysAccount, &jonysAccount)
	fmt.Println(marysAccount)

	//// Pointers
	s1 := CurrentAccount{"P", 1, 10, 1.0}
	s2 := CurrentAccount{"P", 1, 10, 1.0}

	fmt.Println("s1", "and s2", "are:", s1 == s2)

	jonysAccount = new(CurrentAccount)
	fmt.Println(*jonysAccount, jonysAccount, &jonysAccount)

}
