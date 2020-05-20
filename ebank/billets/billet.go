package billets

import "gosamples/ebank/accounts"

/*Pay is*/
func Pay(account accounts.Account, ammount float64) {
	account.Withdraw(ammount)
}
