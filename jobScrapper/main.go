package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50" // 이후 limit=50&start=50, 100, 150 꼴로 올라감

func main() {
	totalPages := getPages() - 1 // 1부터 '다음'까지이므로 찾아볼 페이지는 하나 빼야함

	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50) // strconv.Itoa는 interger to ask 약자. int -> str 형변환
	fmt.Println(pageURL)
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
