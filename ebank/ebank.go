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
	mock2Account := CurrentAccount{"Mock Acks Two", 943, 3476, 1200}
	fmt.Println("Withdraw:", mockAccount.balance, mock2Account.balance)
	mockAccount.Withdraw(200)
	mock2Account.Withdraw(200)
	fmt.Println("Withdraw:", mock2Account.balance, mock2Account.balance)

	// Tests variadics functions
	fmt.Println("Sum:", sum(1, 2, 3))
	sliceNums := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}
	fmt.Println("Sum:", sum(sliceNums...))

	// Tests Deposits
	mock3Account := CurrentAccount{"John Wick Doe", 123, 1234, 5000.}
	mock3Account.Deposit(100)
	fmt.Println(mock3Account)

	// Tests Transfer...
	mock4Account := CurrentAccount{"John Karmac Doe", 313, 8454, 9800.} // He transfers to
	mock5Account := CurrentAccount{"John Lenon Doe", 633, 6534, 6400.}  // Him
	mock4Account.Transfer(450, &mock5Account)
	fmt.Println(mock4Account, mock5Account)
}

// ( c *CurrentAccount is to similar java `this` )
func (c *CurrentAccount) Withdraw(amount float64) float64 {
	canWithDraw := amount >= 0 && amount <= c.balance
	if canWithDraw {
		c.balance -= amount
	}
	return c.balance
}

func (c *CurrentAccount) Deposit(amount float64) float64 {
	canDeposit := amount > 0
	if canDeposit {
		c.balance += amount
	}
	return c.balance
}

func (c *CurrentAccount) Transfer(amount float64, targetAccount *CurrentAccount) bool {
	canTransfer := amount < c.balance && amount > 0
	if canTransfer {
		c.balance -= amount
		targetAccount.Deposit(amount)
		return true
	}
	return false
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
