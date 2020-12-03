package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
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

// Scrape main func
func Scrape(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term + "&limit=50" // 이후 limit=50&start=50, 100, 150 꼴로 올라감
	var jobs []extreactedJob
	totalPages := getPages(baseURL) - 1 // 1부터 '다음'까지이므로 찾아볼 페이지는 하나 빼야함

	mainC := make(chan []extreactedJob)

	// 전체 페이지를 돌면서 jobs를 슬라이스로 받아옴
	for i := 0; i < totalPages; i++ {
		//TODO: 각 페이지 별로 고루틴을 돌리자
		go getPage(i, baseURL, mainC) // 해당 페이지의 job을 struct slice로 만들어 반환함
	}

	// 반환한 채널 합치기
	for i := 0; i < totalPages; i++ {
		extractedJobs := <-mainC
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done", len(jobs), "line")
}

func writeJobs(jobs []extreactedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	// Write any buffered data to the underlying writer (standard output).

	// flsuh와 close의 차이가 무엇이냐? https://stackoverflow.com/a/49166489
	// Closing does cause a flush. When you call Flush and then Close, the stream is flushed twice
	defer w.Flush()

	headers := []string{"ID", "Location", "Title", "Summary"}
	wErr := w.Write(headers)
	checkErr(wErr)

	// TODO: 이 녀석도 고루틴을 돌리고 싶은데 concurrency 문제
	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.location, job.title, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func getPage(page int, url string, mainC chan<- []extreactedJob) {
	var jobs []extreactedJob

	// 채널 생성
	c := make(chan extreactedJob)

	pageURL := url + "&start=" + strconv.Itoa(page*50) // strconv.Itoa는 interger to ask 약자. int -> str 형변환
	fmt.Println("Requesting : " + pageURL)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // memory leak 방지를 위해 close
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find("div.jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		//TODO: card에서 정보 추출하는 과정 고루틴
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extreactedJob) {
	// Attr는 가져온 dom의 속성을 반환함
	id, _ := card.Attr("data-jk")
	location := cleanString(card.Find(".sjcl>span.location").Text())
	title := cleanString(card.Find("h2.title>a").Text())
	summary := cleanString(card.Find(".summary").Text())

	c <- extreactedJob{id, location, title, summary}
}

func cleanString(str string) string {
	// func Join(a []string, sep string) string: 문자열 슬라이스에 저장된 문자열을 모두 연결
	// strings.Fields 함수는 공백을 기준으로 문자열을 쪼개어 문자열 슬라이스로 저장합니다.
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages(url string) int {

	pages := 0
	res, err := http.Get(url)

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
