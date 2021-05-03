package main

import (
	"fmt"

	"github.com/seokyu/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("seokyu")
	fmt.Println(account)
}
