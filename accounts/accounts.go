package accounts

import (
	"errors"
	"fmt"
)

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

var errWithDraw = errors.New("can't Widthdraw")

// Withdraw x amount from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errWithDraw
	}
	a.balance -= amount
	return nil
}

// Change Owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}
