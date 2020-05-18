package accounts

// CurrentAccount is
type CurrentAccount struct {
	Holder  string
	Agency  int
	Number  int
	Balance float64
}

// ( c *CurrentAccount is to similar java `this` )
func (c *CurrentAccount) Withdraw(amount float64) float64 {
	canWithDraw := amount >= 0 && amount <= c.Balance
	if canWithDraw {
		c.Balance -= amount
	}
	return c.Balance
}

func (c *CurrentAccount) Deposit(amount float64) float64 {
	canDeposit := amount > 0
	if canDeposit {
		c.Balance += amount
	}
	return c.Balance
}

func (c *CurrentAccount) Transfer(amount float64, targetAccount *CurrentAccount) bool {
	canTransfer := amount < c.Balance && amount > 0
	if canTransfer {
		c.Balance -= amount
		targetAccount.Deposit(amount)
		return true
	}
	return false
}
