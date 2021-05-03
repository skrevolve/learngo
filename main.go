package main

import (
	"fmt"

	"github.com/seokyu/learngo/banking"
)

func main() {
	account := banking.Account{Owner: "seogyu", Balance: 1000}
	fmt.Println(account)
}
