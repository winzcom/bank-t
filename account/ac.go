package account

type Accounts map[string]int

func CreateAccounts() Accounts {
	m := make(Accounts, 0)
	return m
}
