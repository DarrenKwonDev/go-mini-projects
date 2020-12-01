package main

import (
	"fmt"
	"net/http"
)

type resultResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan resultResult)

	urls := []string{"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://cineps.net/"}

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

func hitURL(url string, c chan resultResult) {
	fmt.Println("checking", url)

	resp, err := http.Get(url)
	status := "OK"

	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		c <- resultResult{url: url, status: "FALED"}
	}
	c <- resultResult{url: url, status: status}
}
