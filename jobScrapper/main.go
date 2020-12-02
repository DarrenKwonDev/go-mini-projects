package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type extreactedJob struct {
	id       string
	location string
	title    string
	salary   string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50" // 이후 limit=50&start=50, 100, 150 꼴로 올라감

func main() {
	totalPages := getPages() - 1 // 1부터 '다음'까지이므로 찾아볼 페이지는 하나 빼야함

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50) // strconv.Itoa는 interger to ask 약자. int -> str 형변환
	fmt.Println("Requesting : " + pageURL)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // memory leak 방지를 위해 close
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find("div.jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		// Attr는 가져온 dom의 속성을 반환함
		id, _ := card.Attr("data-jk")
		fmt.Println(id)
		title := card.Find("h2.title>a").Text()
		location := card.Find(".sjcl>span.location").Text()

		fmt.Println(title)
		fmt.Println(location)

	})

}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)

	// error handling
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // memory leak 방지를 위해 close

	// caustion! It is the responsibility of the caller to close it if required.
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("li").Length() // 1부터 다음까지
	})

	return pages
}

/*
error handling funcs
*/

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err) // kill the program
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("request failed with Status:", res.StatusCode)
	}
}
