package accounts

// Account is
type Account interface {
	Withdraw(amount float64) float64
}
