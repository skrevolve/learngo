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

// package main

// import (
// 	"fmt"

// 	"github.com/seokyu/learngo/mydict"
// )

// func main() {
// 	dictionary := mydict.Dictionary{}
// 	baseWord := "hello"
// 	// definition := "Greeting"
// 	// err := dictionary.Add(word, definition)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// hello, _ := dictionary.Search(word)
// 	// fmt.Println(hello)

// 	// err2 := dictionary.Add(word, definition)
// 	// if err2 != nil {
// 	// 	fmt.Println(err2)
// 	// }
// 	dictionary.Add(baseWord, "first")
// 	// err := dictionary.Update("asdfasdf", "second")
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	dictionary.Search(baseWord)
// 	dictionary.Delete(baseWord)
// 	word, err := dictionary.Search(baseWord)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(word)
// 	}
// 	fmt.Println(word)
// }
// --------------------------------------------

package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errorRequestFailed = errors.New("Request failed")

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://soundcloud.com",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
