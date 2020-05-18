package main

import (
	"fmt"
	"gosamples/ebank/accounts"
)

func main() {
	// testsPointers()
	mockAccount := accounts.CurrentAccount{"John Mock Doe", 343, 3434, 500.}
	mock2Account := accounts.CurrentAccount{"Mock Acks Two", 943, 3476, 1200}
	fmt.Println("Withdraw:", mockAccount.Balance, mock2Account.Balance)
	mockAccount.Withdraw(200)
	mock2Account.Withdraw(200)
	fmt.Println("Withdraw:", mock2Account.Balance, mock2Account.Balance)

	// Tests variadics functions
	fmt.Println("Sum:", sum(1, 2, 3))
	sliceNums := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}
	fmt.Println("Sum:", sum(sliceNums...))

	// Tests Deposits
	mock3Account := accounts.CurrentAccount{"John Wick Doe", 123, 1234, 5000.}
	mock3Account.Deposit(100)
	fmt.Println(mock3Account)

	// Tests Transfer...
	mock4Account := accounts.CurrentAccount{"John Karmac Doe", 313, 8454, 9800.} // He transfers to
	mock5Account := accounts.CurrentAccount{"John Lenon Doe", 633, 6534, 6400.}  // Him
	mock4Account.Transfer(450, &mock5Account)
	fmt.Println(mock4Account, mock5Account)
}

// Variadic functions
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func testsPointers() {
	petersAccount := accounts.CurrentAccount{"John Doe", 9889, 73784, 199909.909}
	leosAccount := accounts.CurrentAccount{"Leo Doe", 343, 3434, 10000000.909}
	fmt.Println(petersAccount)
	fmt.Println(leosAccount)

	// Other way to create a object
	var marysAccount *accounts.CurrentAccount
	var jonysAccount *accounts.CurrentAccount
	marysAccount = new(accounts.CurrentAccount)
	jonysAccount = new(accounts.CurrentAccount)
	marysAccount.Holder = "Mary Doe"
	jonysAccount.Holder = "Jony Bale"

	fmt.Println(*jonysAccount, jonysAccount, &jonysAccount)
	fmt.Println(marysAccount)

	//// Pointers
	s1 := accounts.CurrentAccount{"P", 1, 10, 1.0}
	s2 := accounts.CurrentAccount{"P", 1, 10, 1.0}

	fmt.Println("s1", "and s2", "are:", s1 == s2)

	jonysAccount = new(accounts.CurrentAccount)
	fmt.Println(*jonysAccount, jonysAccount, &jonysAccount)
}
