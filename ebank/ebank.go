package main

import (
	"fmt"
	"gosamples/ebank/accounts"
	"gosamples/ebank/billets"
	"gosamples/ebank/customers"
)

func main() {
	// testsPointers()
	testsAccounts()
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
	perterCustomer := customers.Customer{"Philip Doe", "000.000.000-11", "Mock"}
	leosCustomer := customers.Customer{"Ana Doe", "110.010.010-01", "Mock"}
	petersAccount := mockAccountBuilder(9889, 73784, 199909.909, perterCustomer)
	leosAccount := mockAccountBuilder(343, 3434, 10000000.929, leosCustomer)
	fmt.Println(petersAccount)
	fmt.Println(leosAccount)

	// Other way to create a object
	var marysAccount *accounts.CurrentAccount
	var jonysAccount *accounts.CurrentAccount
	marysAccount = new(accounts.CurrentAccount)
	jonysAccount = new(accounts.CurrentAccount)
	marysAccount.Holder = customers.Customer{"July Doe", "700.070.900-11", "Mock Mary Doe"}
	jonysAccount.Holder = customers.Customer{"Jony Bale", "092.030.080-11", "Mock"}

	fmt.Println(*jonysAccount, jonysAccount, &jonysAccount) // E.g. Out: {{  } 0 0 0} &{{  } 0 0 0} 0xc000102018
	fmt.Println(marysAccount)

	// Pointers
	s1 := mockAccountBuilder(1, 10, 1.0, perterCustomer)
	s2 := mockAccountBuilder(1, 10, 1.0, leosCustomer)

	fmt.Println("s1", "and s2", "are:", s1 == s2)

	jonysAccount = new(accounts.CurrentAccount)
	fmt.Println(*jonysAccount, jonysAccount, &jonysAccount)
}

func testsAccounts() {
	mockAccount := mockAccountBuilder(323, 4234, 2960., customers.Customer{"Mock Acks One", "101.011.001-10", "Mock2"})
	mock2Account := mockAccountBuilder(943, 3476, 1200, customers.Customer{"Mock Acks Two", "001.001.001-01", "Mock2"})
	fmt.Println("Withdraw:", mockAccount.GetBalance(), mock2Account.GetBalance())
	mockAccount.Withdraw(200)
	mock2Account.Withdraw(200)
	fmt.Println("Withdraw:", mock2Account.GetBalance(), mock2Account.GetBalance())

	// Tests variadics functions
	fmt.Println("Sum:", sum(1, 2, 3))
	sliceNums := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233}
	fmt.Println("Sum:", sum(sliceNums...))

	// Tests Deposits
	mock3Account := mockAccountBuilder(123, 1234, 5000., customers.Customer{"John Wick Doe", "100.100.100-10", "Researcher"})
	mock3Account.Deposit(100)
	fmt.Println(mock3Account)

	// Tests Transfer...
	mock4Account := mockAccountBuilder(313, 8454, 1900, customers.Customer{"John Karmac Doe", "010.010.010-01", "Dev"})   // He transfers to
	mock5Account := mockAccountBuilder(633, 6534, 1200, customers.Customer{"John Lenon Doe", "101.010.110-10", "Singer"}) // Him
	mock4Account.Transfer(460, &mock5Account)
	fmt.Println(mock4Account, mock5Account)

	// Tests get balance
	mock6Account := accounts.CurrentAccount{}
	mock6Account.Deposit(100)
	fmt.Println(mock6Account.GetBalance())

	// Tests billet Payments
	mock7Account := mockAccountBuilder(6263, 82384, 48989, customers.Customer{"Josh Lookort", "934.783.039-83", "Doctor"})
	billets.Pay(&mock7Account, 300)
	fmt.Println(mock7Account)

	savingAccount := accounts.SavingAccount{}
	savingAccount.Deposit(600)
	billets.Pay(&savingAccount, 480)

}

func mockAccountBuilder(agency int, number int, ammount float64, customer customers.Customer) accounts.CurrentAccount {
	var acountReference *accounts.CurrentAccount
	acountReference = new(accounts.CurrentAccount)
	acountReference.Holder = customer
	acountReference.Agency = agency
	acountReference.Number = number
	acountReference.Deposit(ammount)
	return *acountReference
}
