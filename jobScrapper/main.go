package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extreactedJob struct {
	id       string
	location string
	title    string
	summary  string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50" // 이후 limit=50&start=50, 100, 150 꼴로 올라감

func main() {
	var jobs []extreactedJob
	totalPages := getPages() - 1 // 1부터 '다음'까지이므로 찾아볼 페이지는 하나 빼야함

	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(i)
		jobs = append(jobs, extractedJobs...)
	}
	fmt.Println(jobs)
}

func getPage(page int) []extreactedJob {
	var jobs []extreactedJob

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
		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return jobs
}

func extractJob(card *goquery.Selection) extreactedJob {
	// Attr는 가져온 dom의 속성을 반환함
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find("h2.title>a").Text())
	location := cleanString(card.Find(".sjcl>span.location").Text())
	summary := cleanString(card.Find(".summary").Text())

	// fmt.Println(id)
	// fmt.Println(title)
	// fmt.Println(location)
	// fmt.Println(summary)
	return extreactedJob{id, title, location, summary}
}

func cleanString(str string) string {
	// func Join(a []string, sep string) string: 문자열 슬라이스에 저장된 문자열을 모두 연결
	// strings.Fields 함수는 공백을 기준으로 문자열을 쪼개어 문자열 슬라이스로 저장합니다.
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
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
