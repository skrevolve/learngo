package accounts

import "fmt"

//Account struct
type Account struct {
	owner   string
	balance int
}

//NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit x account on your account
func (a *Account) Deposit(amount int) {
	fmt.Println("gonna Despoit", amount)
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) {
	a.balance -= amount
}
