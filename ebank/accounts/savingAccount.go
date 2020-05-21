package accounts

import "gosamples/ebank/customers"

// SavingAccount is
type SavingAccount struct {
	Holder                    customers.Customer
	Agency, Operation, Number int
	balance                   float64
}

// Withdraw is
func (c *SavingAccount) Withdraw(amount float64) float64 { // (c *SavingAccount) is to similar to this in anothers langs.
	canWithDraw := amount >= 0 && amount <= c.balance
	if canWithDraw {
		c.balance -= amount
	}
	return c.balance
}

// Deposit is
func (c *SavingAccount) Deposit(amount float64) float64 {
	canDeposit := amount > 0
	if canDeposit {
		c.balance += amount
	}
	return c.balance
}

// GetBalance is
func (c *SavingAccount) GetBalance() float64 {
	return c.balance
}
