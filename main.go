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

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// )

// var errorRequestFailed = errors.New("Request failed")

// type requestResult struct {
// 	url    string
// 	status string
// }

// func main() {
// 	results := make(map[string]string)
// 	c := make(chan requestResult)
// 	urls := []string{
// 		"https://www.airbnb.com",
// 		"https://www.google.com",
// 		"https://www.amazon.com",
// 		"https://www.reddit.com",
// 		"https://www.facebook.com",
// 		"https://www.instagram.com",
// 		"https://soundcloud.com",
// 	}

// 	for _, url := range urls {
// 		go hitURL(url, c)
// 	}

// 	for i := 0; i < len(urls); i++ {
// 		result := <-c
// 		results[result.url] = result.status
// 	}

// 	for url, status := range results {
// 		fmt.Println(url, status)
// 	}
// }

// func hitURL(url string, c chan<- requestResult) {
// 	resp, err := http.Get(url)
// 	status := "OK"
// 	if err != nil || resp.StatusCode >= 400 {
// 		status = "FAILED"
// 	}
// 	c <- requestResult{url: url, status: status}
// }
// -------------------------------------------------------

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery" //you have to download in $GOPATH(/src/github.com/)
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkStatusCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	searchCards := doc.Find(".jobsearch-SerpJobCard ")
	searchCards.Each(func(i int, card *goquery.Selection) {
		id, _ := card.Attr("data-jk")
		title := card.Find(".title>a").Text()
		location := card.Find(".sjcl").Text()
		fmt.Println(id, title, location)
	})
}

func cleanString(str string) string {

}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkStatusCode(res)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status : ", res.StatusCode)
	}
}
