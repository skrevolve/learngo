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
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, _ := dictionary.Search(word)
	// fmt.Println(hello)

	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	dictionary.Add(baseWord, "first")
	err := dictionary.Update(baseWord, "second")
	if err != nil {
		fmt.Println(err)
	}

	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
}
