package main

import (
	"fmt"

	"github.com/seokyu/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("seokyu")
	account.Deposit(10)
	fmt.Println(account.Balance())
	account.Withdraw(10)
	fmt.Println(account.Balance())
}
