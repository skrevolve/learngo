// package main

// import (
// 	"fmt"

// 	"github.com/seokyu/learngo/accounts"
// )

// func main() {
// 	account := accounts.NewAccount("seokyu")
// 	account.Deposit(10)
// 	// err := account.Withdraw(20)
// 	// if err != nil {
// 	// 	//log.Fatalln(err)
// 	// 	fmt.Println(err)
// 	// }
// 	fmt.Println(account)
// }
//----------------------------------------------

package main

import (
	"fmt"

	"github.com/seokyu/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "firstword"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
