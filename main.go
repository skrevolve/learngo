package main

import (
	"fmt"
	"log"

	"github.com/seokyu/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("seokyu")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance(), account.Owner())
}
