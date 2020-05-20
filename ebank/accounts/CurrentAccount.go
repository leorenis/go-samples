package accounts

import "gosamples/ebank/customers"

// CurrentAccount is
type CurrentAccount struct {
	Holder         customers.Customer
	Agency, Number int
	balance        float64
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

func (c *CurrentAccount) GetBalance() float64 {
	return c.balance
}
